import { AuthModule, DatabaseModule, RmqModule } from '@app/common';
import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { MongooseModule } from '@nestjs/mongoose';
import * as Joi from 'joi';
import { QuoteOrdersRepository } from './QuoteOrders.repository';
import { QuoteOrdersController } from './quote-orders.controller';
import { QuoteOrdersService } from './quote-orders.service';
import { Order, QuoteOrdersSchema } from './schemas/order.schema';
import { TICKET_GENERATION_SERVICE } from './constants/services';
 
@Module({
  imports: [ConfigModule.forRoot({
    // ConfigModule now accessible in entire quote-orders with isGlobal:true
    isGlobal: true,
    // validation schema to make sure env is set while mongoDB connection instantiation
    validationSchema: Joi.object({
      MONGODB_URI: Joi.string().required(),
      PORT: Joi.number().required(),
    }),
    // mention file path as each microservice will have its own set of .env
    envFilePath: './apps/quote-orders/.env'
  }),
  // importing DatabaseModule from common libs to make use of common CRUD repository and schema code inside of quote-orders
  DatabaseModule,
  // register the Order Schema
  // with array of objects with each schema
  MongooseModule.forFeature([{name: Order.name,schema: QuoteOrdersSchema}]),
  // 📝 importing dynamic rmq module here so that quote-orders can register the ticket-generation microservice and use it to communicate with it
  // registering the ticket-generation service inside of quote-orders modules via rabbitmq module
  RmqModule.register({
    name: TICKET_GENERATION_SERVICE
  }),
  //importing auth-jwt-guard and AuthModule from common lib to make use of jwt strategy and apply JwtAuthGuard via @Req nestjs/common
  // 🏷 note the AuthModule make use of AuthService -> register()  to register new queue that uses the Auth queue from .env so make sure to set it in .env 
  AuthModule
],
  controllers: [QuoteOrdersController],
  // instantiate the QuoteOrdersServices & QuoteOrdersRepository
  providers: [QuoteOrdersService,QuoteOrdersRepository],
})
export class QuoteOrdersModule {}
