import { createParamDecorator,ExecutionContext } from "@nestjs/common";
import { GqlExecutionContext } from "@nestjs/graphql";

/**
 * @desc usecase to determine current user logged in details on basis of rest or gql request type
 */
export const CurrentUser = createParamDecorator(
    (_data: unknown, context: ExecutionContext) => {
        
        if(context.getType() === 'http'){
            return context.switchToHttp().getRequest().user;
        }

        const ctx = GqlExecutionContext.create(context);
        return ctx.getContext().req.user;
    }
)