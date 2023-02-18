import { AuthGuard } from "@nestjs/passport"

/**
 * @desc- usecase- for any REST-API endpoint 
 * @reff automatically to the jwt.strategy under strategies folder
 * */
export class JwtAuthGuard extends AuthGuard('jwt') {}