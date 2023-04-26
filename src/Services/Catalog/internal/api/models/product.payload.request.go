/*
 * @File: product.payload.req.go
 * @Description: Return payload info for product
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */
package models

type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required,gt=0,lt=255" example:"Iphone 14 Pro"`
	Price       float64 `json:"price" binding:"required,min=0.01" example:"1400"`
	Category    string  `json:"categoryid"`
	Brand       string  `json:"brandid"`
	Summary     string  `json:"summary" binding:"required,gt=0,lt=10000" example:"Iphone 14 Pro Gold"`
	Description string  `json:"description" binding:"required,gt=0,lt=10000" example:"Iphone 14 Pro Gold 256GB"`
	ImageFile   string  `json:"imageFile" binding:"required" example:"default.png"`
	// ProductType string `json:"producttype" binding:"required,gt=0,lt=255" example:"phone"`
}

type RequestUpdateProduct struct {
	Name        string  `json:"name" binding:"required,gt=0,lt=255" example:"Iphone 14 Pro"`
	Price       float64 `json:"price" binding:"required,min=0.01" example:"1400"`
	Category    string  `json:"categoryid"`
	Brand       string  `json:"brandid"`
	Summary     string  `json:"summary" binding:"required,gt=0,lt=10000" example:"Iphone 14 Pro Gold"`
	Description string  `json:"description" binding:"required,gt=0,lt=10000" example:"Iphone 14 Pro Gold 256GB"`
	ImageFile   string  `json:"imageFile" binding:"required" example:"default.png"`
}

type CreateProductImageRequest struct {
	ImageUrl string
	IsMain   bool
}

type UpdateProductRequest struct {
	Name        string  `json:"name" binding:"required,gt=0,lt=255" example:"Iphone 14 Pro"`
	Price       float64 `json:"price" binding:"required,min=0.01" example:"1400"`
	Category    string  `json:"categoryid"`
	Brand       string  `json:"brandid"`
	Summary     string  `json:"summary" binding:"required,gt=0,lt=10000" example:"Iphone 14 Pro Gold"`
	Description string  `json:"description" binding:"required,gt=0,lt=10000" example:"Iphone 14 Pro Gold 256GB"`
	ImageFile   string  `json:"imageFile" binding:"required" example:"default.png"`
}
