import { AuthModule, RmqModule } from '@app/common';
import { Module } from '@nestjs/common';
import * as Joi from 'joi';
import { ConfigModule } from '@nestjs/config';
import { TicketGenerationController } from './ticket-generation.controller';
import { TicketGenerationService } from './ticket-generation.service';
 
@Module({
  // imporring env via ConfigModule as rabbit mq needs env's
  imports: [
    ConfigModule.forRoot({
      isGlobal: true,
      validationSchema: Joi.object({
      RABBIT_MQ_URI: Joi.string().required(),
      RABBIT_MQ_TICKETGENERATION_QUEUE: Joi.string().required(),
      }),
    }),
  RmqModule,
  AuthModule
],
  controllers: [TicketGenerationController],
  providers: [TicketGenerationService],
})
export class TicketGenerationModule {}
