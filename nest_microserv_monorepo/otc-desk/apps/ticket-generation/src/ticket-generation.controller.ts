import { Controller, Get } from '@nestjs/common';
import { TicketGenerationService } from './ticket-generation.service';

@Controller()
export class TicketGenerationController {
  constructor(private readonly ticketGenerationService: TicketGenerationService) {}

  @Get()
  getHello(): string {
    return this.ticketGenerationService.getHello();
  }
}
