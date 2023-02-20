import {
    Injectable,
    NestInterceptor,
    ExecutionContext,
    CallHandler,
    Logger
} from '@nestjs/common';
import { Observable } from 'rxjs';
import { tap } from 'rxjs/operators'
import { v4 as uuidv4 } from 'uuid';

/**@desc app-level logging interceptor that logs all type of request based on request type http,rpc,websocket */
@Injectable()
export class LoggingInterceptor implements NestInterceptor {
    
    private readonly logger = new Logger(LoggingInterceptor.name);

    private logHttpCall(context: ExecutionContext, next: CallHandler){
        const request = context.switchToHttp().getRequest();
        const userAgent = request.get('user-agent') || '';
        const {ip,method,path:url} = request;
        const corelationKey = uuidv4();
        const userId = request.user?.userId;

        this.logger.debug(
            `[${corelationKey}] ${method} ${url} ${userId} ${userAgent} ${ip}: ${context.getClass().name} ${context.getHandler().name}`,
        );
        const now = Date.now();
        return next.handle().pipe(
            tap(()=> {
                const response = context.switchToHttp().getResponse();

                const {statusCode} = response;
                const contentLength = response.get('content-length');

                this.logger.debug(
                    `[${corelationKey}] ${method} ${url} ${statusCode} ${contentLength} ${Date.now()-now}ms`,
                );
            }),
        );
    }

    intercept(context: ExecutionContext, next: CallHandler): Observable<any> {
        if(context.getType() === 'http'){
            return this.logHttpCall(context, next);
        }
    }
}