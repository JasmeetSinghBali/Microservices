import { Injectable } from '@nestjs/common';

@Injectable()
export class TicketGenerationService {
  getHello(): string {
    return 'Hello World!';
  }
}
