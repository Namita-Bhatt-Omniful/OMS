package main

import (
	"OMS/inits"
	"fmt"

	"github.com/omniful/go_commons/http"
)

func main() {
	server := http.InitializeServer(":8081", 0, 0, 0)
	// routes.Test(server)
	inits.InitializeDB()
	inits.InitializeSQS()
	err := server.StartServer("OMS")

	if err != nil {
		fmt.Println("Failed to start the server:", err)
	}

	fmt.Println("Server started successfully!")
}
