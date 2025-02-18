package inits

import (
	"context"
	"fmt"

	"github.com/omniful/go_commons/kafka"
	"github.com/omniful/go_commons/pubsub"
)

var KafkaProducer *kafka.ProducerClient
var KafkaConsumer *kafka.ConsumerClient

// Implement message handler
type MessageHandler struct{}

func (h *MessageHandler) Handle(ctx context.Context, msg *pubsub.Message) error {
	fmt.Println("processing message:", msg)
	return nil
}
func InitializeKafka() {
	// initialize consumer
	consumer := kafka.NewConsumer(
		kafka.WithBrokers([]string{"localhost:9092"}), //sets the no of kafka brokers(broker address)
		kafka.WithConsumerGroup("my-consumer-group"),  //sets group name for kafka consumer group
		kafka.WithClientID("my-consumer"),             //takes the client id to use for creating sarama config
		kafka.WithKafkaVersion("3.8.1"),               //kafka version to use for creating sarama config
	)
	defer consumer.Close()
	KafkaConsumer = consumer

	//initialize producer
	producer := kafka.NewProducer(
		kafka.WithBrokers([]string{"localhost:9092"}),
		kafka.WithClientID("my-producer"),
		kafka.WithKafkaVersion("2.8.1"),
	)
	defer producer.Close()
	KafkaProducer = producer

}
