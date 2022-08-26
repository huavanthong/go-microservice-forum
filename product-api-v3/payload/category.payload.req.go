/*
 * @File: category.payload.req.go
 * @Description: Return payload info for product
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */
package payload

import (
	"github.com/huavanthong/microservice-golang/product-api-v3/models"
)

type RequestCreateCategory struct {
	Name        string             `json:"name" bson:"name" binding:"required,gt=0,lt=255" example:"Iphone 14 Pro"`
	SubCategory models.SubCategory `json:"category" bson:"category" binding:"required,gt=0,lt=255" example:"Phone"`
	Description string             `json:"description" bson:"description" binding:"required,gt=0,lt=10000" example:"Iphone 14 Pro Gold 256GB"`
}

type RequestUpdateCategory struct {
	Name        string             `json:"name" bson:"name" binding:"required,gt=0,lt=255" example:"Iphone 14 Pro"`
	Category    models.SubCategory `json:"category" bson:"category" binding:"required,gt=0,lt=255" example:"Phone"`
	Description string             `json:"description" bson:"description" binding:"required,gt=0,lt=10000" example:"Iphone 14 Pro Gold 256GB"`
}
