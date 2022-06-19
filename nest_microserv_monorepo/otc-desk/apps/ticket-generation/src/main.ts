import { RmqService } from '@app/common';
import { NestFactory } from '@nestjs/core';
import { TicketGenerationModule } from './ticket-generation.module';

async function bootstrap() {
  const app = await NestFactory.create(TicketGenerationModule);
  // bootstraping rabbit mq with ticket-generation
  const rmqService = app.get<RmqService>(RmqService)
  app.connectMicroservice(rmqService.getOptions('TICKETGENERATION'))
  // use startAllMicroservices instead of app.listen
  await app.startAllMicroservices();
}
bootstrap();
