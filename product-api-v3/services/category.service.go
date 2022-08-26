package services

import (
	"github.com/huavanthong/microservice-golang/product-api-v3/models"
	"github.com/huavanthong/microservice-golang/product-api-v3/payload"
)

type CategoryService interface {
	CreateCategory(pr *payload.RequestCreateCategory) (*models.Category, error)
	FindAllCategories(page int, limit int) ([]*models.Category, error)
	FindCategoryByID(id string) (*models.Category, error)
	FindCategoryByName(name string) (*[]models.Category, error)
	UpdateCategory(id string, pr *payload.RequestUpdateCategory) (*models.Category, error)
	DeleteCategory(id string) error
}
