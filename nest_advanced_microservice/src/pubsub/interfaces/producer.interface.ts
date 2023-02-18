/**@desc generic IProducer interafce that every producer have to implement */
export interface IProducer {
    connect: () => Promise<void>
    disconnect: () => Promise<void>
    produce: (message: any) => Promise<void>;
}