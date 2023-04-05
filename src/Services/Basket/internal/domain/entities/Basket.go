package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Basket struct {
	ID             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID         string             `json:"user_id" bson:"user_id"`
	UserName       string             `json:"user_name" bson:"user_name"`
	Items          []BasketItem       `json:"items" bson:"items"`
	TotalPrice     float64            `json:"total_price" bson:"total_price"`
	TotalDiscounts float64            `json:"total_discounts" bson:"total_discounts"`
	CreatedAt      time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at" bson:"updated_at"`
	ExpiresAt      time.Time          `json:"expires_at" bson:"expires_at"`
}
