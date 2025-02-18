package interservice

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/omniful/go_commons/http"

	interservice_client "github.com/omniful/go_commons/interservice-client"
)

var Client *interservice_client.Client

func InterServiceClient() {
	config := interservice_client.Config{
		ServiceName: "user-service",
		BaseURL:     "http://ocalhost:8080/api",
		Timeout:     5 * time.Second,
	}
	client, err := interservice_client.NewClientWithConfig(config)
	if err != nil {
		panic(err)
	}
	Client = client
}
func GetRequest(ctx context.Context, userData interface{}, url string) (interface{}, *interservice_client.Error) {
	request := &http.Request{
		Url: url,
	}
	_, err := Client.Get(request, &userData)
	if err != nil {
		return nil, err
	}
	json_data, _ := json.Marshal(userData)
	fmt.Println(string(json_data))
	return &userData, nil
}

func PostRequest(ctx context.Context, userData interface{}, url string, body interface{}) (interface{}, *interservice_client.Error) {
	request := &http.Request{
		Url:  url,
		Body: body,
	}
	_, err := Client.Post(request, &userData)
	if err != nil {
		return nil, err
	}
	json_data, _ := json.Marshal(userData)
	fmt.Println(string(json_data))
	return &userData, nil
}
