import { IsNotEmpty, IsPhoneNumber, IsPositive, IsString } from "class-validator";

// create order request dto with class validators to check the passed param in request dto for creatin order
export class createOrderRequest{
    @IsString()
    @IsNotEmpty()
    name: string;

    @IsPositive()
    price: number;

    @IsPhoneNumber()
    phoneNumber: string;
}