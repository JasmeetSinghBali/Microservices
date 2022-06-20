import { Injectable, UnauthorizedException } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { PassportStrategy } from '@nestjs/passport';
import { ExtractJwt, Strategy } from 'passport-jwt';
import { Types } from 'mongoose';
import { TokenPayload } from '../auth.service';
import { UsersService } from '../users/users.service';

// This jwt strategy is used to authenticate request and then grab user details ref to this jwt token
@Injectable()
export class JwtStrategy extends PassportStrategy(Strategy) {
  constructor(
    configService: ConfigService,
    private readonly usersService: UsersService,
  ) {
    super({
      // extracting jwt from the request
      jwtFromRequest: ExtractJwt.fromExtractors([
        (request: any) => {
          return request?.Authentication;
        },
      ]),
      secretOrKey: configService.get('JWT_SECRET'),
    });
  }
  
  // get user details in ref to the jwt  
  async validate({ userId }: TokenPayload) {
    try {
      return await this.usersService.getUser({
        _id: new Types.ObjectId(userId),
      });
    } catch (err) {
      throw new UnauthorizedException();
    }
  }
}
