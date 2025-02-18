package services

import (
	"OMS/csv"
	"OMS/inits"
	"OMS/interservice"
	"OMS/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID            string `json:"id"`
	SellerID      string `json:"seller_id"`
	ItemCount     int    `json:"item_count"`
	ModeOfPayment string `json:"mode_of_payment"`
	Status        string `json:"status"`
	Amount        int    `json:"amount"`
}

func CreateBulkOrder(filePath string) {
	// file is the 2d slice having records fetched from csv
	file := csv.ParseCSVFile(filePath)
	var ValidOrders []Order
	userData, _ := interservice.PostRequest(context.Background(), &ValidOrders, "/sku/verify", file)
	fmt.Println(userData)
	collection := inits.DB.Database("oms_db").Collection("orders")
	for _, value := range ValidOrders {
		order := models.Order{
			ID:            primitive.NewObjectID(),
			SellerID:      value.SellerID,
			ItemCount:     value.ItemCount,
			ModeOfPayment: value.ModeOfPayment,
			Status:        value.Status,
			Amount:        value.Amount,
		}
		_, err := collection.InsertOne(context.Background(), order)
		if err != nil {
			fmt.Println("Error in inserting an order in DB : ", err)
		}

	}
}
