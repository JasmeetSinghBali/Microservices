import { Injectable } from '@nestjs/common';
import { createOrderRequest } from './dto/create-order.request';
import { QuoteOrdersRepository } from './QuoteOrders.repository';

@Injectable()
export class QuoteOrdersService {
  constructor(private readonly quoteOrdersRepository: QuoteOrdersRepository){}
  // service that call create orders repos with Orders model to create a new order
  async createOrder(request: createOrderRequest){
    // call the common(abstract) create repository method
    return this.quoteOrdersRepository.create(request);
  }
  // service that call abstract repos with Orders model to retrieve all orders
  async getOrders(){
    return this.quoteOrdersRepository.find({});
  }
}
