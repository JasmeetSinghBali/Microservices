import { Injectable, UnauthorizedException } from "@nestjs/common";
import { PassportStrategy } from "@nestjs/passport";
import { Strategy } from "passport-local";
import { User } from "src/user/models/user";
import { AuthService } from "../auth.service";

@Injectable()
export class LocalStrategy extends PassportStrategy(Strategy){
    constructor (private readonly authService: AuthService){
        super({ usernameField: 'email'})
    }

    async validate(email: string, password: string): Promise<User|boolean> {
        const user = await this.authService.validate(email,password);
        if (!user){
            return false;
        }

        return new Promise<User|boolean>((resolve)=>{
            resolve(user)
        });
    }
}