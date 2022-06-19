import { Module } from '@nestjs/common';
import { TicketGenerationController } from './ticket-generation.controller';
import { TicketGenerationService } from './ticket-generation.service';

@Module({
  imports: [],
  controllers: [TicketGenerationController],
  providers: [TicketGenerationService],
})
export class TicketGenerationModule {}
