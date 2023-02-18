import { Injectable } from "@nestjs/common";
import { InjectConnection } from "@nestjs/mongoose";
import { Connection } from "mongoose";

@Injectable()
export class DatabaseService {
    constructor(@InjectConnection() private readonly connection: Connection){}

    getDbConnection(): Connection {
        return this.connection;
    }
}