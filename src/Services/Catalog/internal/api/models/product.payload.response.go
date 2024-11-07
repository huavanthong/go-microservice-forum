/*
 * @File: product.payload.resp.go
 * @Description: Return payload info for product
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */
package models

import (
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/domain/entities"
)

// Error defines the response error
type CreateProductSuccess struct {
	Status  string `json:"status" example:"success"`
	Code    int    `json:"code" example:"201"`
	Message string `json:"message" example:"Create a new post success"`
	Data    entities.Product
}

type GetAllProductSuccess struct {
	Status  string `json:"status" example:"success"`
	Code    int    `json:"code" example:"201"`
	Message string `json:"message" example:"Get all products success"`
	Data    []*entities.Product
}

type GetProductSuccess struct {
	Status  string `json:"status" example:"success"`
	Code    int    `json:"code" example:"201"`
	Message string `json:"message" example:"Get product success"`
	Data    *entities.Product
}

type GetProductsSuccess struct {
	Status  string `json:"status" example:"success"`
	Code    int    `json:"code" example:"201"`
	Message string `json:"message" example:"Get product success"`
	Data    []*entities.Product
}
type UpdateProductSuccess struct {
	Status  string `json:"status" example:"success"`
	Code    int    `json:"code" example:"201"`
	Message string `json:"message" example:"Update a exist post success"`
	Data    entities.Product
}
