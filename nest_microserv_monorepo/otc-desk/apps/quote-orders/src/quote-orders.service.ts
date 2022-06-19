import { Inject, Injectable } from '@nestjs/common';
import { ClientProxy } from '@nestjs/microservices';
import { TICKET_GENERATION_SERVICE } from './constants/services';
import { createOrderRequest } from './dto/create-order.request';
import { QuoteOrdersRepository } from './QuoteOrders.repository';

@Injectable()
export class QuoteOrdersService {
  // injecting the registered ticket-generation service inside here to communicate with it
  // ref: https://docs.nestjs.com/microservices/rabbitmq
  // clientProxy helps to exchanges messages between microservices
  constructor(private readonly quoteOrdersRepository: QuoteOrdersRepository, @Inject(TICKET_GENERATION_SERVICE) private ticketGen: ClientProxy){}
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
