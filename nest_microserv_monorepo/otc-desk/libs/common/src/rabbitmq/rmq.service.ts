// rabbit-mq injectible services 

import { Injectable } from "@nestjs/common";
import { ConfigService } from "@nestjs/config";
import { RmqOptions, Transport } from "@nestjs/microservices";

// implement rabbit mq functionality
@Injectable()
export class RmqModule{
    constructor (private readonly configService:ConfigService){}
    // COMMON-INITIALIZER-RABBITMQ for microservices
    // takes name of the rabbit mq queue to be initialized by a microservice
    // with noAck = false we need to manually acknowledge the message before removing it from the queue
    // noAck=true by default i.e Nestjs will automatically acknowlege these rabbitmq messages
    getOptions(queue: string, noAck = false): RmqOptions{
        return{
            transport: Transport.RMQ,
            options:{
                // where rabbit mq be listening on
                urls: [this.configService.get<string>('RABBIT_MQ_URI')],
                // name of the rabbit queue for the microservice
                queue: this.configService.get<string>(`RABBIT_MQ_${queue}_QUEUE`),
                noAck,
                persistent: true // so that the queue maintains the list of messages
            }
        }
    }
}