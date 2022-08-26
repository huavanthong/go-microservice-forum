/*
 * @File: product.payload.req.go
 * @Description: Return payload info for product
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */
package payload

import (
	"github.com/huavanthong/microservice-golang/product-api-v3/models"
)

type RequestCreateProduct struct {
	Name        string          `json:"name" bson:"name" binding:"required,gt=0,lt=255" example:"Iphone 14 Pro"`
	Category    models.Category `json:"category" bson:"category" binding:"required,gt=0,lt=255" example:"Phone"`
	Summary     string          `json:"summary" bson:"summary" binding:"required,gt=0,lt=10000" example:"Iphone 14 Pro Gold"`
	Description string          `json:"description" bson:"description" binding:"required,gt=0,lt=10000" example:"Iphone 14 Pro Gold 256GB"`
	ImageFile   string          `json:"imageFile" bson:"imageFile" binding:"required" example:"default.png"`
	Price       float64         `json:"price" bson:"price" binding:"required,min=0.01" example:"1400"`
}

type RequestUpdateProduct struct {
	Name        string          `json:"name" bson:"name" binding:"required,gt=0,lt=255" example:"Iphone 14 Pro"`
	Category    models.Category `json:"category" bson:"category" binding:"required,gt=0,lt=255" example:"Phone"`
	Summary     string          `json:"summary" bson:"summary" binding:"required,gt=0,lt=10000" example:"Iphone 14 Pro Gold"`
	Description string          `json:"description" bson:"description" binding:"required,gt=0,lt=10000" example:"Iphone 14 Pro Gold 256GB"`
	ImageFile   string          `json:"imageFile" bson:"imageFile" binding:"required" example:"default.png"`
	Price       float64         `json:"price" bson:"price" binding:"required,min=0.01" example:"1400"`
}
