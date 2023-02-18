import { Module } from '@nestjs/common';
import { GraphQLModule } from '@nestjs/graphql';
import { UserModule } from './user/user.module';
import { ApolloDriver, ApolloDriverConfig } from '@nestjs/apollo';
import { AuthModule } from './auth/auth.module';
import { ConfigModule } from '@nestjs/config';
import { PubsubModule } from './pubsub/pubsub.module';
import * as Joi from 'joi';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { MockTestKafkaConsumer } from './mock-test.consumer';
import { PrometheusModule } from '@willsoto/nestjs-prometheus/dist/module';
import { APP_INTERCEPTOR } from '@nestjs/core';
import { LoggingInterceptor } from './util/logging.interceptor';

@Module({
  imports: [
    ConfigModule.forRoot({
      isGlobal: true,
      validationSchema: Joi.object({
        JWT_SECRET: Joi.string().required(),
        JWT_EXPIRATION: Joi.string().required(),
        MONGODB_URI: Joi.string().required(),
        KAFKA_ENDPOINT: Joi.string().required(),
        KAFKA_MECHANISM: Joi.string().required(),
        KAFKA_USERNAME: Joi.string().required(),
        KAFKA_PASSWORD: Joi.string().required(),
      }),
      envFilePath: '.env'
    }),
    GraphQLModule.forRoot<ApolloDriverConfig>({
      driver: ApolloDriver,
      autoSchemaFile: true
    }),
    UserModule,
    AuthModule,
    PubsubModule,
    PrometheusModule.register()
  ],
  controllers: [AppController],
  providers: [AppService,MockTestKafkaConsumer, {
    provide: APP_INTERCEPTOR,
    useClass: LoggingInterceptor,
  }],
})
export class AppModule {}
