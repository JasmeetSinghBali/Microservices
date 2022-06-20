import { JwtAuthGuard } from '@app/common';
import { Body, Controller, Get, Post, Req, UseGuards } from '@nestjs/common';
import { createOrderRequest } from './dto/create-order.request';
import { QuoteOrdersService } from './quote-orders.service';

// localhost:3000/orders
@Controller('orders')
export class QuoteOrdersController {
  constructor(private readonly quoteOrdersService: QuoteOrdersService) {}
  
  // create order
  @Post()
  @UseGuards(JwtAuthGuard)
  async createOrder(
    @Body() request: createOrderRequest, @Req() req: any
  ){
    // @Req object to extract the request object
    console.log(req.user);// if JwtAuthGuard succeded then the request object will have user on it
    return this.quoteOrdersService.createOrder(request);
  }
  // get orders
  @Get()
  async getOrders(){
    return this.quoteOrdersService.getOrders();
  }
}
