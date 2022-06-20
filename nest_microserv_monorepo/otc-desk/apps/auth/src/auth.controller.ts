import { Controller, Post, Res, UseGuards } from '@nestjs/common';
import { MessagePattern } from '@nestjs/microservices';
import { Response } from 'express';
import { AuthService } from './auth.service';
import { CurrentUser } from './current-user.decorator';
import JwtAuthGuard from './guards/jwt-auth.guard';
import { LocalAuthGuard } from './guards/local-auth.guard';
import { User } from './users/schemas/user.schema';

@Controller('auth')
export class AuthController {
  constructor(private readonly authService: AuthService) {}
  
  // local auth guard refers to only using email-password to login
  @UseGuards(LocalAuthGuard)
  @Post('login')
  async login(
    @CurrentUser() user: User,
    @Res({ passthrough: true }) response: Response,
  ) {
    await this.authService.login(user, response);
    response.send(user);
  }
  
  // 📝 ClientProxy send request from jwt-auth-guard message pattern
  // Authenticates the current user by grabbing it via context that we recieve from message pattern send as request intially by jwt-auth-guard in libs common via current-user.decorator.ts with rpc|http case mentioned
  @UseGuards(JwtAuthGuard)
  @MessagePattern('validate_user')
  async validateUser(@CurrentUser() user: User){
    // 📝 this will return user back to the calling service if the user is authenticated by jwt auth guard 
    return user;
  }
}
