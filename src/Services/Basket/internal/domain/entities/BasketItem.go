package entities

import "time"

type BasketItem struct {
	ProductID      string    `json:"product_id" bson:"product_id"`
	ProductName    string    `json:"product_name" bson:"product_name"`
	Quantity       int       `json:"quantity" bson:"quantity"`
	Price          float64   `json:"price" bson:"price"`
	DiscountAmount float64   `json:"discount_amount" bson:"discount_amount"`
	TotalPrice     float64   `json:"total_price" bson:"total_price"`
	ImageURL       string    `json:"image_url" bson:"image_url"`
	CreatedAt      time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" bson:"updated_at"`
}
