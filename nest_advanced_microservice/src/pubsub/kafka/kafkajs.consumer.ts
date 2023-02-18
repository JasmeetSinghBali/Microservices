import { Logger } from "@nestjs/common";
import { Consumer, ConsumerConfig, ConsumerSubscribeTopics, Kafka, KafkaMessage} from "kafkajs";
import { sleep } from "src/util/sleep";
import { IConsumer } from "../interfaces/consumer.interface";
import * as retry from 'async-retry';
import { DatabaseService } from "src/database/database.service";

/**@desc specific implementation of the ConsumerInterface IConsumer */
export class KafkajsConsumer implements IConsumer{
    private readonly kafka: Kafka; // 💡 connection to broker
    private readonly consumer: Consumer; // 💡 consumer instance
    private readonly logger: Logger; // 💡 logger specific to a consumer

    /**@desc initial values surrounding this consumer instance */
    constructor(
        private readonly topic: ConsumerSubscribeTopics,
        private readonly databaseService: DatabaseService,
        config: ConsumerConfig,
        broker: string
    ){
        this.kafka = new Kafka({
            brokers: [broker],
            sasl: {
                mechanism: process.env.KAFKA_MECHANISM as any,
                username: process.env.KAFKA_USERNAME as string,
                password: process.env.KAFKA_PASSWORD as string,
              },
            ssl: true,
        });
        this.consumer = this.kafka.consumer(config)
        this.logger = new Logger(`${topic.topics}-${config.groupId}`);   
    }

    /**@desc repeated failed tries reconnects by consumer to specific broker after 5 seconds*/
    async connect() {
        try{
            // 💡 kafkajs will try 5 times, if still fails then reaches the catch blocks
            await this.consumer.connect();
        }catch(err: any){
            this.logger.error('failed to connect consumer<>broker reff: kafkajs.consumer',err);
            // reconnection delay while retrying to connect to broker
            await sleep(5000);
            await this.consumer.connect();
        }
    }

    /**@desc pushes failed messages to dead letter queue by interacting with MongoDB */
    private async addMessageToDeadLetterQueue(message: KafkaMessage){
        await this.databaseService.getDbConnection().collection('deadlq').insertOne({value: message.value, topic: this.topic.topics});
    }

    /**@desc absracted consume method implementation with abstracted eachMessage */
    async consume(onMessage: (message: KafkaMessage) => Promise<void>) {
        await this.consumer.subscribe(this.topic);
        await this.consumer.run({
            // 💡 below code runs on each message consumed, in case failed consumer will crash and try to run again
            eachMessage: async ({message,partition})=>{
                this.logger.debug(`Processing message partition ${partition}`);
                // 💡 in case of failed consumption, retry 2 more times if still fails then push the message to the dead letter queue to persist in MongoDB
                try{
                    // keeps on retrying i.e calling the onMessage function
                    // 📝 on  2 consecutive fails this goes to catch block
                    await retry(async () => onMessage(message), {
                        retries: 2,
                        onRetry: (error,attempt) => 
                            this.logger.error(
                                `failed to consume message, executing retry ${attempt}/2`, error
                            )
                    });
                }catch(err: any){
                    this.logger.error(`failed to consume message from topic: ${JSON.stringify(this.topic)} adding to dead-letter-queue...`,err);
                    await this.addMessageToDeadLetterQueue(message);
                }
            },
        });
    }

    async disconnect() {
        await this.consumer.disconnect();
    } 
}