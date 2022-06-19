import { Injectable, Logger } from '@nestjs/common';

@Injectable()
export class TicketGenerationService {
  // logger that is specific to this ticket-generation service
  private readonly logger = new Logger(TicketGenerationService.name);
  getHello(): string {
    return 'Hello World!';
  }
  generateTicket(data: any){
    this.logger.log('Creating ticket on the basis of data that we recieved while placing order...',data);
  }
}
