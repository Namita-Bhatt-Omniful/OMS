package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"` //_id is default unique identifier in mongodb
	SellerID      string             `bson:"seller_id,omitempty" json:"seller_id"`
	ItemCount     int                `bson:"item_count" json:"item_count"`
	ModeOfPayment string             `bson:"mode_of_payment" json:"mode_of_payment"`
	Amount        int                `bson:"amount" json:"amount"`
	Status        string             `bson:"status" json:"status"`
	Address       string             `bson:"address" json:"address"`
	CreatedAt     primitive.DateTime `bson:"created_at" json:"created_at"`
	UpdatedAt     primitive.DateTime `bson:"updated_at" json:"updated_at"`
}
