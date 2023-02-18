import { InputType, Field } from '@nestjs/graphql';
import { IsNotEmpty, IsOptional } from 'class-validator';

@InputType()
export class UpdateUserInput {
    @Field()
    @IsNotEmpty()
    user_ID: string;

    // 💡 the IsNotEmpty wont be triggered if the email key in the mutation payload is not provided
    @Field()
    @IsOptional()
    @IsNotEmpty()
    email?: string;
}