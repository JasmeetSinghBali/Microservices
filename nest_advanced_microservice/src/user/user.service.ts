import { Injectable } from "@nestjs/common";
import { GetUserArgs } from "./dto/args/get-user.args";
import { GetUsersArgs } from "./dto/args/get-users.args";
import { CreateUserInput } from "./dto/input/create-user.input";
import { UpdateUserInput } from "./dto/input/update-user.input";
import { User } from "./models/user";

@Injectable()
export class UserService {
    
    // Usecase: Public via resolvers

    public async createUser(createUserData: CreateUserInput): Promise<User | boolean> {
        // 🎈abstract-repository call goes here
        return false;
    }
    
    public async updateUser(updateUserData: UpdateUserInput): Promise<User | boolean> {
        // 🎈abstract-repository call goes here
        return false;
    }
    
    public async getUser(getUserArgs: GetUserArgs): Promise<User | null> {
        // 🎈abstract-repository call goes here
        return null;
    }

    public async getUsers(getUsersArgs: GetUsersArgs): Promise<User[] | null> {
        // 🎈abstract-repository call goes here
        return null;
    }

    // Usecase: Internal methods
    public async getUserByEmail(email: string): Promise<User | null> {
        // 🎈 abstract-repository call goes here
        return null;
    }
}