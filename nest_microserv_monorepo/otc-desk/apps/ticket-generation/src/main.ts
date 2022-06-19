import { NestFactory } from '@nestjs/core';
import { TicketGenerationModule } from './ticket-generation.module';

async function bootstrap() {
  const app = await NestFactory.create(TicketGenerationModule);
  await app.listen(3000);
}
bootstrap();
