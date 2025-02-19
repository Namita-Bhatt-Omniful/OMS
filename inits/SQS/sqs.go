package sqs

import (
	"OMS/config"
	order "OMS/services/BulkOrder"
	"context"
	"fmt"

	"github.com/omniful/go_commons/sqs"
)

var Queue *sqs.Queue
var SQSConsumer *sqs.Consumer
var SQSPublisher *sqs.Publisher

type MyHandler struct{}

// func (h *MyHandler) Handle(mssg *sqs.Message) error {
// 	fmt.Println("Handling message:", string(mssg.Value))
// 	services.CreateBulkOrder(string(mssg.Value))
// 	return nil
// }

// a message handler that implements the ISqsMessageHandler interface
func (h *MyHandler) Process(ctx context.Context, mssgs *[]sqs.Message) error {
	fmt.Println("processing sqs message")
	for _, mssg := range *mssgs {
		order.CreateBulkOrder(string(mssg.Value))
	}
	return nil
}

func InitializeSQS() {
	// create a queue
	queue_name := "OMS_SQS.fifo"
	q, err := sqs.NewFifoQueue(context.Background(), queue_name, config.SQSconfig)
	if err != nil || q == nil {
		fmt.Println("Queue Initialization error", err)
		return
	}
	Queue = q
	// create consumer
	Handler := &MyHandler{}
	//handler is of sqs.ISqsMessageHandler type which is an interface with Process method inside it
	//I could pass MyHandler type here as it is implementing Process method
	consumer, err := sqs.NewConsumer(q, uint64(1), 1, Handler, 5, 30, true, false)
	if err != nil {
		fmt.Println("Error in creating consumer", err)
		return
	}
	SQSConsumer = consumer
	consumer.Start(context.Background())
	publisher := sqs.NewPublisher(q)
	SQSPublisher = publisher
	fmt.Println("SQS Initialization successful!")

}
