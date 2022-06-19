import { Body, Controller, Get, Post } from '@nestjs/common';
import { createOrderRequest } from './dto/create-order.request';
import { QuoteOrdersService } from './quote-orders.service';

// localhost:3000/orders
@Controller('orders')
export class QuoteOrdersController {
  constructor(private readonly quoteOrdersService: QuoteOrdersService) {}
  
  // create order
  @Post()
  async createOrder(
    @Body() request: createOrderRequest
  ){
    return this.quoteOrdersService.createOrder(request);
  }
  // get orders
  @Get()
  async getOrders(
    @Body() request: createOrderRequest
  ){
    return this.quoteOrdersService.getOrders();
  }
}
