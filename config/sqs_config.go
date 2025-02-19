package config

import "github.com/omniful/go_commons/sqs"

var SQSconfig = &sqs.Config{
	Account:  "779846812549",
	Endpoint: "https://sqs.eu-north-1.amazonaws.com/779846812549/OMS_SQS.fifo",
	Region:   "eu-north-1",
}
