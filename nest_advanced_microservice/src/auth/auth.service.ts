import { Injectable, Logger } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { User } from 'src/user/models/user';
import { UserService } from 'src/user/user.service';

@Injectable()
export class AuthService {
    
    private readonly logger = new Logger(AuthService.name);

    constructor(
        private readonly userService: UserService,
        private readonly jwtService: JwtService,
    ){}

    async validate(email: string, password: string): Promise<User | null>{
        // 🎈 abstract repository call to check user exists in DB
        const user = await this.userService.getUserByEmail(email);
        if(!user){
            return null
        }

        // 🎈 compare password
        return null
    }

    async login(user: User): Promise<{ access_token: string }> {
        const payload = {
            email: user.email,
            sub: user.user_ID
        }

        return new Promise<{access_token: string}>((resolve)=>{
            resolve(Object.freeze({
                access_token: this.jwtService.sign(payload),
            }))
        }); 
    }

    async verify(token: string): Promise<User|boolean> {
        const decoded = this.jwtService.verify(token, {
            secret: process.env.JWT_SECRET as string
        })
        const user = await this.userService.getUserByEmail(decoded.email);
        if(!user){
            this.logger.error('failed to grab user from decoded token payload')
            return false
        }
        return user;
    }
}
