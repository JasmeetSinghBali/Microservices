import { ValidationPipe } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { NestFactory } from '@nestjs/core';
import { QuoteOrdersModule } from './quote-orders.module';

async function bootstrap() {
  const app = await NestFactory.create(QuoteOrdersModule);
  // it makes sure all the endpoints are protected from receiving incorrect data
  app.useGlobalPipes(new ValidationPipe());
  const configService = app.get(ConfigService)
  await app.listen(configService.get('PORT'));
}
bootstrap();
