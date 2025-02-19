package worker

import (
	"fmt"

	"github.com/omniful/go_commons/pubsub"
)

func KafkaConsumer(msg *pubsub.Message) {

	fmt.Println("processing message:", msg)
}
