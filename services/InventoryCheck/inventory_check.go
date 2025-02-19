package inventorycheck

import (
	dbconn "OMS/inits/DB"
	"OMS/interservice"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

type Order struct {
	OrderID       string `json:"id"`
	SellerID      string `json:"seller_id"`
	SKUID         string `json:"sku_id"`
	ItemCount     string `json:"item_count"`
	ModeOfPayment string `json:"mode_of_payment"`
	Status        string `json:"status"`
}
type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

// message value is of []byte type
func ValidateInventory(bytedata []byte) {
	var order Order
	json.Unmarshal(bytedata, &order) //parses the json-encoded message value into order
	var response Response
	_, err := interservice.PostRequest(context.Background(), &response, "/inventory/update", order)
	if err != nil {
		fmt.Println("Error:", err)
	}
	collection := dbconn.DB.Database("oms_db").Collection("orders")
	//if inventory check is okay,update order status
	// updating status corresponding to the order id of valid orders
	if response.Status == string(http.StatusOK) {
		filter := bson.M{"OrderID": order.OrderID}
		update := bson.M{"$set": bson.M{"status": "new_order"}}

		_, err := collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			fmt.Println(err)
		}
	}
}
