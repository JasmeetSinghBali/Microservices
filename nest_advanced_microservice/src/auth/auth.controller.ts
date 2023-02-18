import { Controller, Post, Req, UseGuards } from "@nestjs/common";
import { AuthService } from "./auth.service";
import { LocalAuthGuard } from "./guards/local-auth.guard";
import { Request } from 'express';
import { User } from "src/user/models/user";

/**
 * @desc usecase to generate json web token JWT against valid username-password pair via AuthService
 */
@Controller('auth')
export class AuthController {
    constructor(private readonly authService: AuthService){}

    /**
     * @desc login uses LocalAuthGuard with local.strategy that attaches user context with the request on validation of username-password
     */
    @UseGuards(LocalAuthGuard)
    @Post('login')
    async login (@Req() req: Request): Promise<{access_token: string}> {
        const {access_token} = await this.authService.login(req.user as User)
        return new Promise<{access_token: string}>((resolve)=>{
            resolve({access_token})
        });
    }
}