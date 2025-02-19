package main

import (
	dbconn "OMS/inits/DB"
	afka "OMS/inits/Kafka"
	"OMS/interservice"
	"OMS/routes"
	"fmt"

	"github.com/omniful/go_commons/http"
)

func main() {
	server := http.InitializeServer(":8081", 0, 0, 0)
	// routes.Test(server)
	routes.GetRouter(server)
	dbconn.InitializeDB()
	interservice.InterServiceClient()
	// Sqs.InitializeSQS()
	afka.InitializeKafka()
	err := server.StartServer("OMS")

	if err != nil {
		fmt.Println("Failed to start the server:", err)
	}

	fmt.Println("Server started successfully!")
}
