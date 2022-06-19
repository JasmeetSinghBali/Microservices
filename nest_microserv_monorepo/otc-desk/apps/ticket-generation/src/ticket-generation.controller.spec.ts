import { Test, TestingModule } from '@nestjs/testing';
import { TicketGenerationController } from './ticket-generation.controller';
import { TicketGenerationService } from './ticket-generation.service';

describe('TicketGenerationController', () => {
  let ticketGenerationController: TicketGenerationController;

  beforeEach(async () => {
    const app: TestingModule = await Test.createTestingModule({
      controllers: [TicketGenerationController],
      providers: [TicketGenerationService],
    }).compile();

    ticketGenerationController = app.get<TicketGenerationController>(TicketGenerationController);
  });

  describe('root', () => {
    it('should return "Hello World!"', () => {
      expect(ticketGenerationController.getHello()).toBe('Hello World!');
    });
  });
});
