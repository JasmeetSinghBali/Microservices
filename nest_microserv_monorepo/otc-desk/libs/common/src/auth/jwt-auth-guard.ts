import { CanActivate, ExecutionContext, Inject, Injectable, UnauthorizedException } from "@nestjs/common";
import { ClientProxy } from "@nestjs/microservices";
import { catchError, Observable } from "rxjs";
import { AUTH_SERVICE } from "./services";

@Injectable()
export class JwtAuthGuard implements CanActivate{
    // 📝  injecting the auth service via inject token to be used in JwtAuthGuard and 
    // 📝  clientProxy exposes send method that is send message to the microservice and return observable with its response
    constructor(@Inject(AUTH_SERVICE) private authClient: ClientProxy){}

    // 📝  we can figure out wheather request is http or rabbitmq request with execution context given to us by nestjs
    canActivate(context: ExecutionContext): boolean | Promise<boolean> | Observable<boolean> {
        const authentication = this.getAuthentication(context);
        // and sends a request to rabbit mq microservice with authentication jwt and wait for response unlike event
        return this.authClient.send('validate_user',{
            Authentication: authentication,
        }).pipe(
            // pipe this rxjs observables with tap operator for side effects
            // to get the response from auth service and then addUser to the current request to data object depending upon rpc or http
            tap((res)=>{
                this.addUser(res,context)
            }),
            // rxjs operator that catch any error in the observable chain and rethrow the exception
            catchError(()=>{
                throw new UnauthorizedException()
            }),
        );
    }

    // 📝 rpc|http find executionContext and a/cgly grab the authentication jwt from cookies or Authentication in case of rabbitmq rpc
    private getAuthentication(context: ExecutionContext){
        let authentication: string;
        // IMP: 📝  in case of rabbit mq the context type would be rpc, and in http the context would be http
        if(context.getType() === 'rpc'){
            //we grab the rpc data and then the Authentication key that holds the jwt
            authentication = context.switchToRpc().getData().Authentication
        }else if(context.getType() === 'http'){
            // grabbing jwt from cookies in case of http and the cookie parser is already adding cookies object to request
            authentication = context.switchToHttp().getRequest().cookies?.Authentication
        }
        if(!authentication){
            throw new UnauthorizedException('No value(rabbitmq[rpc-authentication]|http[jwt-cookies-parser]) was provided for authentication');
        }
        return authentication;
    }

    // adds user to the request http or rpc
    private addUser(user:any,context: ExecutionContext){

    }
}