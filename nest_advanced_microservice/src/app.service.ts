import { Injectable } from '@nestjs/common';
import { ProducerService } from './pubsub/service/producer.service';
import * as fibonacci from 'fibonacci';

@Injectable()
export class AppService {
  constructor(
    private readonly producerService: ProducerService,
  ) {}
  hello(){
    return 'hello world'
  }
  /**@desc mock kafka producer message */
  async testKafkaProducer() {
    await this.producerService.produce(
      'wb-tickets',
      { 
        value: 'this is message from kafka-producer mock-test-service in app.service'
      }
    );
    return 'succesfully produced test message to kafka for topic- wb-tickets';
  }

  /** 🎈 MAKE SURE TO COMMENT THIS FLOW OUT DONT KEEP IT FOR PRODUCTION, also remove fibonacci package
   * @desc cpu intensive mock service call */
  stressMockCall(): number{
    return fibonacci.iterate(1000);
  }
}