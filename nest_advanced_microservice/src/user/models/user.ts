import { ObjectType, Field } from '@nestjs/graphql';

@ObjectType()
export class User {
    @Field()
    user_ID: string;
    @Field()
    email: string;
    @Field({ nullable: true })
    password?: string 
}