import { Injectable, Logger, OnModuleInit } from "@nestjs/common";
import { ConsumerService } from "./pubsub/service/consumer.service";

@Injectable()
export class MockTestKafkaConsumer implements OnModuleInit {
    
    private readonly logger = new Logger(MockTestKafkaConsumer.name);

    constructor(private readonly consumerService: ConsumerService){}
    
    /**@desc setup consumer listening to the mock-test */
    async onModuleInit() {
        await this.consumerService.consume(
            {
                topic: { topics: ['wb-tickets'] },
                config: { groupId: 'test-consumer' },
                onMessage: async(message) => {
                    this.logger.debug(message.value.toString());
                    throw new Error('This is an intentional test error that simulates a failed consumed message to check dead letter queue setup!');
                },
            }
        );
    }
}