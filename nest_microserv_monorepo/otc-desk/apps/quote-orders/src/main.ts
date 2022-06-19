import { NestFactory } from '@nestjs/core';
import { QuoteOrdersModule } from './quote-orders.module';

async function bootstrap() {
  const app = await NestFactory.create(QuoteOrdersModule);
  await app.listen(3000);
}
bootstrap();
