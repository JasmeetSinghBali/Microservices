import { Injectable } from "@nestjs/common";
import { PassportStrategy } from "@nestjs/passport";
import { ExtractJwt, Strategy } from "passport-jwt";
import { User } from "src/user/models/user";
import { UserService } from "src/user/user.service";

@Injectable()
export class JwtStrategy extends PassportStrategy(Strategy){
    constructor(
        private readonly userService: UserService
    ){
        super({
            jwtFromRequest: ExtractJwt.fromAuthHeaderAsBearerToken(),
            ignoreExpiration: false,
            secretOrKey: process.env.JWT_SECRET as string
        })
    }

    async validate(validationPayload: {email: string, sub: string}): Promise<User | boolean> {
        const user = await this.userService.getUserByEmail(validationPayload.email);
        if(!user){
            return false
        }
        return new Promise<User|null>((resolve)=>{
            resolve(user)
        })
    }
}