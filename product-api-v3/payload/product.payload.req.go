/*
 * @File: product.payload.req.go
 * @Description: Return payload info for product
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */
package payload

type RequestCreateProduct struct {
	Name        string  `json:"name" binding:"required,gt=0,lt=255" example:"Iphone 14 Pro"`
	ProductType string  `json:"producttype" binding:"required,gt=0,lt=255"`
	Category    string  `json:"category" binding:"required,gt=0,lt=255"`
	Summary     string  `json:"summary" binding:"required,gt=0,lt=10000" example:"Iphone 14 Pro Gold"`
	Description string  `json:"description" binding:"required,gt=0,lt=10000" example:"Iphone 14 Pro Gold 256GB"`
	ImageFile   string  `json:"imageFile" binding:"required" example:"default.png"`
	Price       float64 `json:"price" binding:"required,min=0.01" example:"1400"`
}

type RequestUpdateProduct struct {
	Name        string  `json:"name" bson:"name" binding:"required,gt=0,lt=255" example:"Iphone 14 Pro"`
	Category    string  `json:"category" bson:"category" binding:"required,gt=0,lt=255"`
	Summary     string  `json:"summary" bson:"summary" binding:"required,gt=0,lt=10000" example:"Iphone 14 Pro Gold"`
	Description string  `json:"description" bson:"description" binding:"required,gt=0,lt=10000" example:"Iphone 14 Pro Gold 256GB"`
	ImageFile   string  `json:"imageFile" bson:"imageFile" binding:"required" example:"default.png"`
	Price       float64 `json:"price" bson:"price" binding:"required,min=0.01" example:"1400"`
}
