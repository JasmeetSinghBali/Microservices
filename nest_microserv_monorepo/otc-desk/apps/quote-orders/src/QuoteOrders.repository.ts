import { AbstractRepository } from "@app/common";
import { Injectable, Logger } from "@nestjs/common";
import { InjectConnection, InjectModel } from "@nestjs/mongoose";
import { Connection, Model } from "mongoose";
import { Order } from "./schemas/order.schema";

@Injectable()
// OrderRepository extends the parent AbstractRepository with CRUD passing the Order Schema
export class QuoteOrdersRepository extends AbstractRepository<Order>{
    // to have a logger that logs the name of the order
    protected readonly logger = new Logger(QuoteOrdersRepository.name)
    // injecting the Order.name , orderModel(label), connection(label) into AbstractRepository via constructor for tracking the database transactions
    constructor(
        @InjectModel(Order.name) orderModel: Model<Order>,
        @InjectConnection() connection: Connection
        ){  
            // pass the two i.e injectedModel & injectedConnection to the parent i.e AbstractRepository via their label
            super(orderModel,connection);
            // which in return allow us to have the CRUD methods define in AbstractRepository
        }
}