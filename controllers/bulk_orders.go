package controllers

import (
	Sqs "OMS/inits/SQS"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omniful/go_commons/sqs"
)

func CreateBulkOrder(ctx *gin.Context) {
	var request struct {
		//get file path from the body of post request
		Address string `json:"address"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// create a message
	data, _ := json.Marshal(request)
	message := &sqs.Message{
		GroupId:         "group-1",
		Value:           data,
		ReceiptHandle:   "group-1",
		DeduplicationId: "gp-1",
	}
	// publish the message to sqs queue
	err := Sqs.SQSPublisher.Publish(context.Background(), message)
	if err != nil {
		fmt.Println("Failed to publish error", err)
		return
	}
	fmt.Println("Message published successfully!")
	// order.CreateBulkOrder(request.Address)
	ctx.JSON(http.StatusOK, gin.H{"message": "success"}) // send success response
}
