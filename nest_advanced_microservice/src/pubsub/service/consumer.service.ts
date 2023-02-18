import { Injectable, OnApplicationShutdown } from "@nestjs/common";
import { DatabaseService } from "src/database/database.service";
import { IConsumer } from "../interfaces/consumer.interface";
import { KafkajsConsumerOptions } from "../kafka/kafkajs-consumer-options.interface";
import { KafkajsConsumer } from "../kafka/kafkajs.consumer";

/**@desc generic consumer service, can import specific ConsumerInterfaces like KafkajsConsumer of kafka */
@Injectable()
export class ConsumerService implements OnApplicationShutdown{
    /**@desc Generic consumers list with type IConsumer */
    private readonly consumers: IConsumer[] = [];
    
    constructor(
        private readonly databaseService: DatabaseService
    ){}

    async consume({topic, config, onMessage}: KafkajsConsumerOptions){
        const consumer = new KafkajsConsumer(
            topic,
            this.databaseService, // passes databaseService instance to kafkajs.consumer
            config,
            process.env.KAFKA_ENDPOINT as string,
        )
        await consumer.connect();
        await consumer.consume(onMessage);
        this.consumers.push(consumer);
    }

    /**@desc disconnects all consumers when OnApplicationShutdown lifecycle hook nestjs is triggered */
    async onApplicationShutdown() {
        for (const consumer of this.consumers){
            await consumer.disconnect();
        }
    }
}