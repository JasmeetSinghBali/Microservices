import { Inject, Injectable } from '@nestjs/common';
import { ClientProxy } from '@nestjs/microservices';
import { lastValueFrom } from 'rxjs';
import { TICKET_GENERATION_SERVICE } from './constants/services';
import { createOrderRequest } from './dto/create-order.request';
import { QuoteOrdersRepository } from './QuoteOrders.repository';

@Injectable()
export class QuoteOrdersService {
  // injecting the registered ticket-generation service inside here to communicate with it
  // ref: https://docs.nestjs.com/microservices/rabbitmq
  // clientProxy helps to exchanges messages between microservices
  constructor(private readonly quoteOrdersRepository: QuoteOrdersRepository, @Inject(TICKET_GENERATION_SERVICE) private ticketGenerationClient: ClientProxy){}
  // service that call create orders repos with Orders model to create a new order
  async createOrder(request: createOrderRequest){
    // database transac session
    const session = await this.quoteOrdersRepository.startTransaction();
    try{
      // call the common(abstract) create repository method
      // 📝 pass session object as second param during placing order
      const order = await this.quoteOrdersRepository.create(request,{session});

      // 📝 emit the created order event to ticket-generation
      // 📝 Note-ClientProxy returns an observable || convert the order creation as promise to handle this via rxjs
      
      // emit the create order event to ticket-generation microservice named as order_created and data of event as the request received by createOrder 
      await lastValueFrom(this.ticketGenerationClient.emit('order_created',{
        request,
      }),);

      // ✔ if we reach here then the event was successfully emitted to ticket-generation microservice
      // finally commit the transaction i.e now the database call will be executed persisting the new order in the database
      await session.commitTransaction();
      return order;
    }catch(err: any){
      // aborts the current tranasction & cancel all database calls
      // if say the quote-orders was not able to communicate with ticket-generation microservice 
      await session.abortTransaction();
      throw err;
    }
  }
  // service that call abstract repos with Orders model to retrieve all orders
  async getOrders(){
    return this.quoteOrdersRepository.find({});
  }
}
