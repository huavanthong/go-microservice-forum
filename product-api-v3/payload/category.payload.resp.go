/*
 * @File: category.payload.resp.go
 * @Description: Return payload info for category
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */
package payload

import "github.com/huavanthong/microservice-golang/product-api-v3/models"

// Error defines the response error
type CreateCategorySuccess struct {
	Status  string `json:"status" example:"success"`
	Code    int    `json:"code" example:"201"`
	Message string `json:"message" example:"Create a new post success"`
	Data    *models.Category
}

type GetAllCategoriesSuccess struct {
	Status  string `json:"status" example:"success"`
	Code    int    `json:"code" example:"201"`
	Message string `json:"message" example:"Get all categories success"`
	Data    []*models.Category
}

type GetCategorySuccess struct {
	Status  string `json:"status" example:"success"`
	Code    int    `json:"code" example:"201"`
	Message string `json:"message" example:"Get category success"`
	Data    *models.Category
}

type GetCategoriesSuccess struct {
	Status  string `json:"status" example:"success"`
	Code    int    `json:"code" example:"201"`
	Message string `json:"message" example:"Get category success"`
	Data    []*models.Category
}

type UpdateCategorySuccess struct {
	Status  string `json:"status" example:"success"`
	Code    int    `json:"code" example:"201"`
	Message string `json:"message" example:"Update a exist category success"`
	Data    *models.Category
}
