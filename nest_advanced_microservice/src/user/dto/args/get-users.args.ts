import { ArgsType, Field } from '@nestjs/graphql';
import { IsArray } from 'class-validator';

@ArgsType()
export class GetUsersArgs {
    @Field(()=>[String])
    @IsArray()
    user_IDS: string[];
}