import { RmqService } from "@app/common";
import { ValidationPipe } from "@nestjs/common";
import { ConfigService } from "@nestjs/config";
import { NestFactory } from "@nestjs/core";
import { RmqOptions } from "@nestjs/microservices";
import { AuthModule } from "./auth.module";

// this is how you setup hybrid application microservice
// that listen on http
// and also on the rabbit mq define on line 16
async function bootstrap() {
  const app = await NestFactory.create(AuthModule)
  const rmqService = app.get<RmqService>(RmqService);
  // connect the rabbitmq AUTH named queue to the auth microservice
  // here we are setting noAck: true by just passing true as we dont need to manually acknowledge the messages as this is req-resp cycle not event based
  app.connectMicroservice<RmqOptions>(rmqService.getOptions('AUTH',true));
  // pass global pipe Validation pipe so to make use of class validator
  app.useGlobalPipes(new ValidationPipe());
  // grab the ConfigService as well to use .env throughout the auth microservice
  const configService = app.get(ConfigService);
  // start all microservices
  await app.startAllMicroservices();
  // and then finally starting the auth microservice on http 3001 
  app.listen(configService.get('PORT'))
}
bootstrap();
