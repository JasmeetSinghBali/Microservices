import { ConsumerConfig, ConsumerSubscribeTopics, KafkaMessage } from "kafkajs";

/**@desc specific kafkajs consumer interface IConsumer */
export interface KafkajsConsumerOptions{
    topic: ConsumerSubscribeTopics,
    config: ConsumerConfig,
    onMessage: (message: KafkaMessage) => Promise<void>
}