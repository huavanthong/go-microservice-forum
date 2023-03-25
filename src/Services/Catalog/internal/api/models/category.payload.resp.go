/*
 * @File: category.payload.resp.go
 * @Description: Return payload info for category
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */
package models

import (
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/domain/entities"
)

// Error defines the response error
type CreateCategorySuccess struct {
	Status  string `json:"status" example:"success"`
	Code    int    `json:"code" example:"201"`
	Message string `json:"message" example:"Create a new post success"`
	Data    *entities.Category
}

type GetAllCategoriesSuccess struct {
	Status  string `json:"status" example:"success"`
	Code    int    `json:"code" example:"201"`
	Message string `json:"message" example:"Get all categories success"`
	Data    []*entities.Category
}

type GetCategorySuccess struct {
	Status  string `json:"status" example:"success"`
	Code    int    `json:"code" example:"201"`
	Message string `json:"message" example:"Get category success"`
	Data    *entities.Category
}

type GetCategoriesSuccess struct {
	Status  string `json:"status" example:"success"`
	Code    int    `json:"code" example:"201"`
	Message string `json:"message" example:"Get category success"`
	Data    []*entities.Category
}

type UpdateCategorySuccess struct {
	Status  string `json:"status" example:"success"`
	Code    int    `json:"code" example:"201"`
	Message string `json:"message" example:"Update a exist category success"`
	Data    *entities.Category
}
