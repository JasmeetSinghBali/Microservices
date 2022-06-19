import { AbstractDocument } from "@app/common";
import { Schema,SchemaFactory,Prop } from "@nestjs/mongoose"

@Schema({ versionKey: false}) // make versionKey:false to avoid handling this
// inherit common libs abstract schema _id prop
export class Order extends AbstractDocument{
    @Prop()
    name: string;

    @Prop()
    price: number;

    @Prop()
    phoneNumber: string;
    
}

// add the Order class Schema and export as OrderSchema
export const QuoteOrdersSchema = SchemaFactory.createForClass(Order)