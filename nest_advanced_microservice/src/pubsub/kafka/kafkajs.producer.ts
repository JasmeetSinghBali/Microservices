import { Logger } from "@nestjs/common";
import { Kafka, Message, Producer } from "kafkajs";
import { sleep } from "src/util/sleep";
import { IProducer } from "../interfaces/producer.interface";

export class KafkajsProducer implements IProducer{
    private readonly kafka: Kafka; // 💡 connection to broker
    private readonly producer: Producer; // 💡 producer instance
    private readonly logger: Logger; // 💡 logger specific to a producer

    constructor(
        private readonly topic: string,
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
        this.producer = this.kafka.producer();
        this.logger = new Logger(topic); 
    }

    async produce(message: Message){
        await this.producer.send({topic: this.topic, messages: [message] });
    }
    
    /**@desc repeated failed tries reconnects by producer to specific broker after 5 seconds*/
    async connect() {
        try{
            // 💡 kafkajs will try 5 times, if still fails then reaches the catch blocks
            await this.producer.connect();
        }catch(err: any){
            this.logger.error('failed to connect producer<>broker reff: kafkajs.producer',err);
            // reconnection delay while retrying to connect to broker
            await sleep(5000);
            await this.producer.connect();
        }
    }

    async disconnect(){
        await this.producer.disconnect();
    }
}