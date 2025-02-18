package inits

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DB *mongo.Client

func InitializeDB() {
	uri := "mongodb://localhost:27017"
	// cancel parent context after 10 seconds,and creates a new context(basically for refreshing)
	//  if DB is taking too long to respond,cancel the request
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	// create a connection with mongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		fmt.Println("Error while connecting with database", err)
		return
	}
	DB = client
	// to verify the connection,even after successful creation of client database may not be responsive
	//it confirms that DB is accepting queries
	err1 := client.Ping(ctx, readpref.Primary())
	if err1 != nil {
		fmt.Println("Error while communicating with database", err)
		return
	}
	fmt.Println("MongoDB connected successfully!")

}
