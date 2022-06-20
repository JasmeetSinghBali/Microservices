import { RmqService } from '@app/common';
import { Controller, Get } from '@nestjs/common';
import { Ctx, EventPattern, Payload, RmqContext } from '@nestjs/microservices';
import { TicketGenerationService } from './ticket-generation.service';

@Controller()
export class TicketGenerationController {
  constructor(
    private readonly ticketGenerationService: TicketGenerationService,
    //injecting RmqService for accessing manual ack() method 
    private readonly rmqService:RmqService
    ) {}

  @Get()
  getHello(): string {
    return this.ticketGenerationService.getHello();
  }

  // to listen to event order_created by quotes-order microservice
  @EventPattern('order_created')
  // grabbing the payload sent by quotes-order and context of the event received in our case rabbit mq context
  async handleOrderCreate(@Payload() data: any, @Ctx() context: RmqContext){
    // generating the ticket with ref to the data received from quote-order endpoint
    this.ticketGenerationService.generateTicket(data);
    // once the ticket is generated successfully
    // manually acknowledge the current message & remove it of the rabbit queue to avoid re-screening this same message multiple times
    this.rmqService.ack(context);
  }
}
