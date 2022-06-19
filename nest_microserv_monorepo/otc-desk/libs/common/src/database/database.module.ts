/*
=================Common DB connection=========================
All the important changes/updates/fixes need to be documented here.

PURPOSE: wire up mongoDB connection and export as common database module
DATED: 19/06/2022
===========================================================
*/

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