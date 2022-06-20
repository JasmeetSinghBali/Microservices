import { AuthGuard } from '@nestjs/passport';

// 📝 JwtAuthGuard implements the jwt strategy
export default class JwtAuthGuard extends AuthGuard('jwt') {}
