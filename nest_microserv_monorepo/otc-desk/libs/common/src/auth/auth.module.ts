import { MiddlewareConsumer, Module, NestModule } from "@nestjs/common";
import cookieParser from "cookie-parser";
import { RmqModule } from "../rabbitmq/rmq.module";
import { AUTH_SERVICE } from "./services";

// whichever other module that implements the AuthModule it will implement cookie parsing
// cookie parsing will help to grab JWT from cookies
@Module({
    // register rmqservice so that  the auth service can comm with rabbit mq service
    imports:[RmqModule.register({name: AUTH_SERVICE})],
    // rexport the RmqModule further with attached AuthModule from here so that any other module that uses AuthModule will have access to AUTH_SERVICE
    exports: [RmqModule],
})
export class AuthModule implements NestModule{
    // inherit from NestModule apply configure method to apply middleware
    configure(consumer: MiddlewareConsumer) {
        // this will grab cookie and append with the current request
        // and we want this for all route in the system so forRoutes(*)
        consumer.apply(cookieParser()).forRoutes('*');
    }
}