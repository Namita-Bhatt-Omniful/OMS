package order

import (
	"OMS/csv"
	dbconn "OMS/inits/DB"
	k "OMS/inits/Kafka"
	"OMS/interservice"
	"OMS/models"
	"context"
	"encoding/json"
	"fmt"

	"github.com/omniful/go_commons/pubsub"
)

type Order struct {
	OrderID       string `json:"id"`
	SellerID      string `json:"seller_id"`
	SKUID         string `json:"sku_id"`
	ItemCount     string `json:"item_count"`
	ModeOfPayment string `json:"mode_of_payment"`
	Status        string `json:"status"`
}

// for checking inventory
type Req struct {
	OrderID   string
	SKUID     string
	ItemCount string
}

type OrderResponse struct {
	ValidOrders []Order `json:"valid_orders"`
}

func CreateBulkOrder(filePath string) {
	data := csv.ParseCSVFile(filePath)
	fmt.Println("Success")
	var resp OrderResponse
	// post csv file at wms service,there validate the sku id and then receive Validorders from there
	_, err1 := interservice.PostRequest(context.Background(), &resp, "/sku/verify", data)
	if err1 != nil {
		fmt.Println("Error ", err1)
	}
	// fmt.Println(resp)
	collection := dbconn.DB.Database("oms_db").Collection("orders")
	// insert valid order one by one in DB
	for _, value := range resp.ValidOrders {
		order := models.Order{
			OrderID:       value.OrderID,
			SellerID:      value.SellerID,
			SKUID:         value.SKUID,
			ItemCount:     value.ItemCount,
			ModeOfPayment: value.ModeOfPayment,
			Status:        "on Hold",
		}
		_, err := collection.InsertOne(context.Background(), order)
		if err != nil {
			fmt.Println("Error in inserting an order in DB : ", err)
		}
		//  publish valid orders to kafka and then validate inventory

		// Create message with key for FIFO ordering

		req := &Req{
			OrderID:   value.OrderID,
			SKUID:     value.SKUID,
			ItemCount: value.ItemCount,
		}
		data, _ := json.Marshal(req)
		mssg := &pubsub.Message{
			Topic: "test-topic",
			// Key is crucial for maintaining FIFO ordering
			// Messages with the same key will be delivered to the same partition in order
			// orderid will be unique for each order
			Key:   order.OrderID,
			Value: data,
			Headers: map[string]string{
				"custom-header": "value",
			},
		}

		// Context with request ID
		type contextKey string
		key := contextKey("request_id")
		ctx := context.WithValue(context.Background(), key, "req-123")
		// type of KafkaProducer implements Publish method
		err1 := k.KafkaProducer.Publish(ctx, mssg)
		if err != nil {
			panic(err1)
		}
		// fmt.Println(string(mssg.Value))
	}
	fmt.Println("Successfully published valid orders to kafka")
}
