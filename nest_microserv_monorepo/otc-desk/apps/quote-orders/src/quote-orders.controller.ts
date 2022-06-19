import { Controller, Get } from '@nestjs/common';
import { QuoteOrdersService } from './quote-orders.service';

@Controller()
export class QuoteOrdersController {
  constructor(private readonly quoteOrdersService: QuoteOrdersService) {}

  @Get()
  getHello(): string {
    return this.quoteOrdersService.getHello();
  }
}
