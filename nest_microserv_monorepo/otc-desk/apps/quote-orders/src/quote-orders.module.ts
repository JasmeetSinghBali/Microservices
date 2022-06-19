import { DatabaseModule } from '@app/common';
import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { MongooseModule } from '@nestjs/mongoose';
import * as Joi from 'joi';
import { QuoteOrdersRepository } from './QuoteOrders.repository';
import { QuoteOrdersController } from './quote-orders.controller';
import { QuoteOrdersService } from './quote-orders.service';
import { Order, QuoteOrdersSchema } from './schemas/order.schema';
 
@Module({
  imports: [ConfigModule.forRoot({
    // ConfigModule now accessible in entire quote-orders
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
  MongooseModule.forFeature([{name: Order.name,schema: QuoteOrdersSchema}])
],
  controllers: [QuoteOrdersController],
  // instantiate the QuoteOrdersServices & QuoteOrdersRepository
  providers: [QuoteOrdersService,QuoteOrdersRepository],
})
export class QuoteOrdersModule {}
