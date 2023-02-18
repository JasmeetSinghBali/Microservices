import { Controller, Get } from "@nestjs/common";
import { AppService } from "./app.service";

@Controller()
export class AppController {
    constructor(private readonly appService: AppService){}
    @Get()
    hello(){
        return this.appService.hello();
    }  
    @Get('pingkafka')
    pingKafkaProducer(){
        return this.appService.testKafkaProducer();
    }


    /**🎈 MAKE SURE TO REMOVE OR COMMENT THIS ROUTE and remove the fibonacci package also
     * @desc just for mocking stress on server to spin up another server instance and test the autoscale setup with k8s */
    @Get('stress')
    stressMockService(): number{
        return this.appService.stressMockCall();
    }
}