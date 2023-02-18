import { Module } from '@nestjs/common';
import { DatabaseModule } from 'src/database';
import { ConsumerService } from './service/consumer.service';
import { ProducerService } from './service/producer.service';

@Module({
    imports: [DatabaseModule],
    providers: [ProducerService,ConsumerService],
    exports: [ProducerService,ConsumerService] // 💡 message can be produced/consumed by the module that import this ProducerService/ConsumerService, reff: mock-test.consumer and app.controller
})
export class PubsubModule {}
