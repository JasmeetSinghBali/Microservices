import { Test, TestingModule } from '@nestjs/testing';
import { QuoteOrdersController } from './quote-orders.controller';
import { QuoteOrdersService } from './quote-orders.service';

describe('QuoteOrdersController', () => {
  let quoteOrdersController: QuoteOrdersController;

  beforeEach(async () => {
    const app: TestingModule = await Test.createTestingModule({
      controllers: [QuoteOrdersController],
      providers: [QuoteOrdersService],
    }).compile();

    quoteOrdersController = app.get<QuoteOrdersController>(QuoteOrdersController);
  });

  describe('root', () => {
    it('should return "Hello World!"', () => {
      expect(quoteOrdersController.getHello()).toBe('Hello World!');
    });
  });
});
