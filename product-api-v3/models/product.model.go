package models

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/************************ Define structure product ************************/

// ErrProductNotFound is an error raised when a product can not be found in the database
var ErrProductNotFound = fmt.Errorf("Product not found")

type Product struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Name        string             `json:"name" bson:"name" binding:"required,gt=0,lt=255"`
	Category    string             `json:"category" bson:"category" binding:"required,gt=0,lt=255"`
	Summary     string             `json:"summary" bson:"summary" binding:"required,gt=0,lt=10000"`
	Description string             `json:"description" bson:"description" binding:"required,gt=0,lt=10000"`
	ImageFile   string             `json:"imageFile" bson:"imageFile" binding:"required"`
	Price       string             `json:"price" bson:"price" binding:"required,min=0.01"`
	SKU         string             `json:"sku" bson:"sku" binding:"required"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

func FilteredResponse(product *Product) Product {
	return Product{
		ID:          product.ID,
		Name:        product.Name,
		Category:    product.Category,
		Summary:     product.Summary,
		Description: product.Description,
		ImageFile:   product.ImageFile,
		Price:       product.Price,
		SKU:         product.SKU,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}
