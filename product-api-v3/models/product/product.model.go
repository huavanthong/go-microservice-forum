package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/************************ Define structure product ************************/
type product struct {
	ID          primitive.ObjectID `bson:"_id" json:"id" example:"5bbdadf782ebac06a695a8e7"`
	ProductCode string             `bson:"pcode" json:"pcode" binding:"required" example:"p123456789"`
	Name        string             `json:"name" bson:"name" binding:"required,gt=0,lt=255" example:"Iphone 14 Pro"`
	Category    string             `json:"category" bson:"category" binding:"required,gt=0,lt=255" example:"Phone"`
	Summary     string             `json:"summary" bson:"summary" binding:"required,gt=0,lt=10000" example:"Iphone 14 Pro Gold"`
	Description string             `json:"description" bson:"description" binding:"required,gt=0,lt=10000" example:"Iphone 14 Pro Gold 256GB"`
	ImageFile   string             `json:"imageFile" bson:"imageFile" binding:"required" example:"default.png"`
	Price       float64            `json:"price" bson:"price" binding:"required,min=0.01" example:"1400"`
	SKU         string             `json:"sku" bson:"sku" example:"ABC-XYZ-OXY"`
	CreatedAt   string             `json:"created_at" bson:"created_at"`
	UpdatedAt   string             `json:"updated_at" bson:"updated_at"`
	DeleteAt    string             `json:"deleted_at" bson:"deleted_at"`
}

/************************ Implement design pattern for Product ************************/

/************************ Filter info ************************/
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
