import { Module } from '@nestjs/common';
import { QuoteOrdersController } from './quote-orders.controller';
import { QuoteOrdersService } from './quote-orders.service';

@Module({
  imports: [],
  controllers: [QuoteOrdersController],
  providers: [QuoteOrdersService],
})
export class QuoteOrdersModule {}
