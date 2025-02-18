package inits

import (
	"OMS/config"
	"context"
	"fmt"

	"github.com/omniful/go_commons/sqs"
)

var Queue *sqs.Queue
var SQSConsumer *sqs.Consumer
var SQSPublisher *sqs.Publisher

type MyHandler struct{}

// a message handler that implements the ISqsMessageHandler interface
func (h *MyHandler) Handle(mssg *sqs.Message) error {
	fmt.Println("Handling message:", string(mssg.Value))
	return nil
}

func (h *MyHandler) Process(ctx context.Context, mssgs *[]sqs.Message) error {
	for _, mssg := range *mssgs {
		err := h.Handle(&mssg)
		if err != nil {
			return err
		}
	}
	return nil
}

func InitializeSQS() {
	// create a queue
	queue_name := "OMS_SQS"
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
	consumer, err := sqs.NewConsumer(q, 1, 2, Handler, 5, 30, true, false)
	if err != nil {
		fmt.Println("Error in creating consumer", err)
		return
	}
	SQSConsumer = consumer
	publisher := sqs.NewPublisher(q)
	SQSPublisher = publisher

}
