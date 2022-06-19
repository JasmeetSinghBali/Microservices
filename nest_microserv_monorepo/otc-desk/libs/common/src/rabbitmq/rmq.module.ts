// rabbit-mq module

import { DynamicModule, Module } from "@nestjs/common";
import { ConfigService } from "@nestjs/config";
import { ClientsModule, Transport } from "@nestjs/microservices";
import { RmqService } from "./rmq.service";

interface RmqModuleOptions {
    name: string;
}

@Module({
    providers: [RmqService],
    exports: [RmqService],
})
// 📝 making this dynamic module by making a static register method accessible by all instance of this class
export class RmqModule{
    static register({name}: RmqModuleOptions): DynamicModule{
        // return a dynamic module
        return {
            module: RmqModule,
            // register the rabbit mq service  via ClientsModule from nestjs/microservice with reff to  name received in register static method arguments.
            imports: [
                // registerAsyn takes array of services(object) we intend to define
                ClientsModule.registerAsync([
                    {
                        name,
                        // using useFactory to inject the ConfigService(needed for instantiating rabbit mq queue)
                        useFactory: (configService: ConfigService) => ({
                            transport: Transport.RMQ,
                            // passing the required options need to instantiate the rabbit mq queue
                            options: {
                                // array of urls on which rabbitmq will be listening on
                                urls: [configService.get<string>('RABBIT_MQ_URI')],
                                // actual name of the queue that we are trying to register
                                queue: configService.get<string>(`RABBIT_MQ_${name}_QUEUE`),
                            },
                        }),
                        inject: [ConfigService],
                    },
                ]),
            ],
            // re-exporting the ClientsModule so that the consuming module i.e which imports it have access to the ClientsModule along with the registered rabbit mq  service that is created with this dynamic module
            exports: [ClientsModule]
        }
    }
}