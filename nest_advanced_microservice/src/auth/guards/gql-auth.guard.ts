import { ExecutionContext } from "@nestjs/common";
import { GqlExecutionContext } from "@nestjs/graphql";
import { AuthGuard } from "@nestjs/passport";

/**
 * @desc usecase for any graphql query or mutation
 * @reff automatically to the jwt.strategy under strategies folder 
 * */
export class GqlAuthGuard extends AuthGuard('jwt'){
    getRequest(context: ExecutionContext): any {
        const ctx = GqlExecutionContext.create(context);
        // 💡 passing back the current request object in graphql context derived from nestjs execution context
        return ctx.getContext().req;
    }
}