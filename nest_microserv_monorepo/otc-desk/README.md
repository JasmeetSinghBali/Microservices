> # Walkthrough setting up Microservices monorepo in Nest.js

> This walkthrough includes tech..

- Nestjs
- mongodb
- docker
- rabbitmq

> project structure

- [x] otc-desk(monorepo)
  - [x] quote-orders - (root)
    - [x] ticket-generation - microservice
    - [x] auth - microservice

> ## Step-1 Creating a monorepo steps

- mkdir otc-desk and cd into it
- nest new orders
- cd to orders
- nest generate app quote-orders

**Inspect nest-cli.json it will have monorepo true json key:value and a common tsconfig & package.json for both orders & quote-orders**

> ## Step-2 cleaning up

- delete the orders folder from the apps folder inside your otc_desk
- remove orders from nest cli under project section
- and change the root & tsconfigpath values as apps/quote-orders

> ## Step-3 creating further microservices

- cd to otc-desk

                # syntax
                nest g app newmicroserviceName

- nest g app ticket-generation
- nest g app auth

> ## Step-4 starting up d/f microservice or root service

                # cd to otc_desk

                # run root service quote-orders
                npm run start:dev

                # run ticket-generation microservice
                npm run start:dev ticket-generation

> ## Step-5 Sharing code between our d/f microservices/apps we need to create a common library

                # be in otc_desk

                # syntax
                nest g library nameOfLibrary

                # in this project case
                nest g library common

                accept the default @app when terminal asks for it

- a new libs folder will be created that has code that can be shared by all microservices
- check nest-cli u will have now common as liberary their under projects section

> ## Step-6 setting up mongoDB in nestjs

- npm i @nestjs/mongoose mongoose (this will now be accessible to all apps & liberaries)
- go to libs-> common -> src -> database-> database.module.ts

              # add under the @Module section
              import { Module } from '@nestjs/common';
              import { ConfigService } from '@nestjs/config';
              import { MongooseModule } from '@nestjs/mongoose';

              @Module({
                imports: [
                  MongooseModule.forRootAsync({
                    useFactory: (configService: ConfigService) => ({
                      uri: configService.get<string>('MONGODB_URI'),
                    }),
                    inject: [ConfigService],
                  }),
                ],
              })

              export class DatabaseModule {}

- forRoot takes the connection URI to connect to mongoDB database
- configService to fetch mongodbURI from env
- forRootAsync to avoid blocking the event loop
- injecting ConfigService to be used by DatabaseModule

              libs-> common -> src
              database
                - abstract.repository
                - abstract.schema
                - database.module

- be in root otc_desk then -> npm i @nestjs/config

> **NOTE- make sure u delete the .git from otc_desk then do git add . git commit**
