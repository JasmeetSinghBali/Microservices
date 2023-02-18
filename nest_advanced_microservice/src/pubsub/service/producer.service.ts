import {Injectable, OnApplicationShutdown} from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { Message } from 'kafkajs';
import { KafkajsProducer } from '../kafka/kafkajs.producer';
import { IProducer } from '../interfaces/producer.interface';

@Injectable()
export class ProducerService implements OnApplicationShutdown{

    constructor(private readonly configService: ConfigService){}

    /**💡 producers map-list of all producer in the app where key is the topic i.e string and value the generic IProducer */
    private readonly producers = new Map<string,IProducer>();

    /**@desc checks producer instance existance in relation to the topic passed */
    private async getProducer(topic: string){
        let producer = this.producers.get(topic);
        // 💡if producer do not exists, create new
        if(!producer) {
            producer = new KafkajsProducer(
                topic, 
                this.configService.get('KAFKA_ENDPOINT')
            );
            await producer.connect();
            this.producers.set(topic,producer);
        }
        // producer variable shud now be populated if was not initially populated
        return producer;

    }
    async produce(topic: string, message: Message){
        const producer = await this.getProducer(topic);
        await producer.produce(message);
    }

    /**@desc take cares of terminating connection b/w kafka-instance<>server */
    async onApplicationShutdown(signal?: string) {
        for (const producer of this.producers.values()){
            await producer.disconnect()
        }
    }

}