import { Logger, UseGuards } from '@nestjs/common';
import  { Resolver, Query, Args, Mutation } from '@nestjs/graphql'
import { CurrentUser } from 'src/auth/decorator/current-user.decorator';
import { GqlAuthGuard } from 'src/auth/guards/gql-auth.guard';
import { GetUserArgs } from './dto/args/get-user.args';
import { GetUsersArgs } from './dto/args/get-users.args';
import { CreateUserInput } from './dto/input/create-user.input';
import { UpdateUserInput } from './dto/input/update-user.input';
import { User } from './models/user'
import { UserService } from './user.service'

@Resolver(() => User)
export class UserResolver {
    
    private readonly logger = new Logger(UserResolver.name);

    constructor(private readonly userService:UserService){}
    
    /**
     * @desc GqlAuth guard that now requires bearer jwt token in auth header before processing the fetchUser request
     * sample- query headers shud have { "Authorization": "Bearer tokenFromRestAPI /auth/login"}
     * meta 💡 @CurretUser is custom decorator to grab current user from context based on rest or gql request type
     */
    @Query(() => User, {name: 'fetchUser', nullable: true} )
    @UseGuards(GqlAuthGuard)
    async getUser( @CurrentUser() user: User, @Args() getUserArgs: GetUserArgs): Promise<User|null> {
        this.logger.debug(`This is from CurrentUser Custom Decorator : ${user}`);
        const userInfo = await this.userService.getUser(getUserArgs);
        return new Promise<User|null>((resolve)=>{
            resolve(userInfo)
        });
    }

    // 💡 nullable: 'items' make sure that items in the list can be null but the Users list cannot be null , if want to allow both null then set to itemsAndList
    @Query(()=> [User],{name:'fetchUsers',nullable:'items'})
    async getUsers(@Args() getUsersArgs: GetUsersArgs): Promise<User[] | null> {
        const userList = await this.userService.getUsers(getUsersArgs);
        return new Promise<User[]|null>((resolve)=>{
            resolve(userList)
        });
    }

    @Mutation(() => User)
    async createUser (@Args('createUserData') createUserData: CreateUserInput): Promise<User|boolean> {
        const newUser = await this.userService.createUser(createUserData);
        return new Promise<User|boolean>((resolve)=>{
            resolve(newUser)
        });
    }

    @Mutation(() => User)
    async updateUser(@Args('updateUserData') updateUserData: UpdateUserInput): Promise<User|boolean>{
        const updateUser = await this.userService.updateUser(updateUserData);
        return new Promise<User|boolean>((resolve)=>{
            resolve(updateUser)
        })
    }

}