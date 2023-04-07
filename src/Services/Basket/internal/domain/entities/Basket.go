package entities

import (
	"time"

	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/ValueObjects"
)

type Basket struct {
	ID             ValueObjects.BasketID `json:"id" bson:"_id,omitempty"`
	UserID         string                `json:"user_id" bson:"user_id"`
	UserName       string                `json:"user_name" bson:"user_name"`
	Items          []BasketItem          `json:"items" bson:"items"`
	TotalPrice     float64               `json:"total_price" bson:"total_price"`
	TotalDiscounts float64               `json:"total_discounts" bson:"total_discounts"`
	CreatedAt      time.Time             `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time             `json:"updated_at" bson:"updated_at"`
	ExpiresAt      time.Time             `json:"expires_at" bson:"expires_at"`
}

func (basket *Basket) AddItem(basketItem BasketItem) (*Basket, error) {

	basket.Items = append(basket.Items, basketItem)

	return basket, nil
}

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

// Value to Object
type BasketItems []BasketItem
