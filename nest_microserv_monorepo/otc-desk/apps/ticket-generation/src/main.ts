import { RmqService } from '@app/common';
import { NestFactory } from '@nestjs/core';
import { TicketGenerationModule } from './ticket-generation.module';

async function bootstrap() {
  const app = await NestFactory.create(TicketGenerationModule);
  const rmqService = app.get<RmqService>(RmqService)
  app.connectMicroservice(rmqService.getOptions('TICKETGENERATION'))
  await app.startAllMicroservices();
}
bootstrap();
