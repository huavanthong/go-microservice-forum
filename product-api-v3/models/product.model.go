package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/************************ Define structure product ************************/
type Product struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	ProductCode string             `bson:"pcode" json:"pcode" binding:"required"`
	Name        string             `json:"name" bson:"name" binding:"required,gt=0,lt=255"`
	Category    string             `json:"category" bson:"category" binding:"required,gt=0,lt=255"`
	Summary     string             `json:"summary" bson:"summary" binding:"required,gt=0,lt=10000"`
	Description string             `json:"description" bson:"description" binding:"required,gt=0,lt=10000"`
	ImageFile   string             `json:"imageFile" bson:"imageFile" binding:"required"`
	Price       float64            `json:"price" bson:"price" binding:"required,min=0.01"`
	SKU         string             `json:"sku" bson:"sku"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

func FilteredResponse(product *Product) Product {
	return Product{
		ID:          product.ID,
		ProductCode: product.ProductCode,
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
