import { AuthGuard } from "@nestjs/passport";

/**
 * @desc usecase to obtain access token payload by passing username-password pair 
 * @reff automatically to the local.strategy in strategies
 * */
export class LocalAuthGuard extends AuthGuard('local'){}