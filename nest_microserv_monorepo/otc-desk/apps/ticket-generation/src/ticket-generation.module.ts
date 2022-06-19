import { RmqModule } from '@app/common';
import { Module } from '@nestjs/common';
import { TicketGenerationController } from './ticket-generation.controller';
import { TicketGenerationService } from './ticket-generation.service';

@Module({
  imports: [RmqModule],
  controllers: [TicketGenerationController],
  providers: [TicketGenerationService],
})
export class TicketGenerationModule {}
