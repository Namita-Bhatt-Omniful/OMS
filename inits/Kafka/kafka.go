package Kafka

import (
	inventorycheck "OMS/services/InventoryCheck"
	"context"
	"fmt"

	"github.com/omniful/go_commons/kafka"
	"github.com/omniful/go_commons/pubsub"
)

var KafkaProducer *kafka.ProducerClient
var KafkaConsumer *kafka.ConsumerClient

// Implement message handler
type MessageHandler struct{}

func (h *MessageHandler) Process(ctx context.Context, msg *pubsub.Message) error {
	fmt.Println("Processing kafka mssg")
	inventorycheck.ValidateInventory(msg.Value)

	return nil
}
func InitializeKafka() {

	//initialize producer
	producer := kafka.NewProducer(
		kafka.WithBrokers([]string{"localhost:9092"}),
		kafka.WithClientID("my-producer"),
		kafka.WithKafkaVersion("3.8.1"),
	)
	// defer producer.Close()
	KafkaProducer = producer

	// initialize consumer
	consumer := kafka.NewConsumer(
		kafka.WithBrokers([]string{"localhost:9092"}), //sets the no of kafka brokers(broker address)
		kafka.WithConsumerGroup("my-consumer-group"),  //sets group name for kafka consumer group
		kafka.WithClientID("my-consumer"),             //takes the client id to use for creating sarama config
		kafka.WithKafkaVersion("3.8.1"),               //kafka version to use for creating sarama config
	)
	// defer consumer.Close()
	KafkaConsumer = consumer

	// Register message handler for topic
	handler := &MessageHandler{}
	consumer.RegisterHandler("my-topic", handler)

	// Start consuming messages
	ctx := context.Background()

	go consumer.Subscribe(ctx)
	fmt.Println("Kafka Initialization successful!")

}
