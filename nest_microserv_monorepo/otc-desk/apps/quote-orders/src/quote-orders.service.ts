import { Injectable } from '@nestjs/common';

@Injectable()
export class QuoteOrdersService {
  getHello(): string {
    return 'Hello World!';
  }
}
