/*
 * @File: product.payload.resp.go
 * @Description: Return payload info for product
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */
package payload

import "github.com/huavanthong/microservice-golang/product-api-v3/models"

// Error defines the response error
type CreateProductSuccess struct {
	Status  string `json:"status" example:"success"`
	Code    int    `json:"code" example:"201"`
	Message string `json:"message" example:"Create a new post success"`
	Data    models.Products
}

type GetAllProductSuccess struct {
	Status  string `json:"status" example:"success"`
	Code    int    `json:"code" example:"201"`
	Message string `json:"message" example:"Get all products success"`
	Data    []models.Products
}

type GetProductSuccess struct {
	Status  string `json:"status" example:"success"`
	Code    int    `json:"code" example:"201"`
	Message string `json:"message" example:"Get product success"`
	Data    []models.Products
}

type UpdateProductSuccess struct {
	Status  string `json:"status" example:"success"`
	Code    int    `json:"code" example:"201"`
	Message string `json:"message" example:"Update a exist post success"`
	Data    models.Products
}
