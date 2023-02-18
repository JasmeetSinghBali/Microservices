/**@desc commmon consumer interface that all consumers implements*/
export interface IConsumer {
    connect: () => Promise<void>; // connects to brokers
    disconnect: () => Promise<void>; // disconnects to brokers
    consume: (onMessage: (message: any) => Promise<void>) => Promise<void> // executes anonymous arrow function that returns Promise of type void
}