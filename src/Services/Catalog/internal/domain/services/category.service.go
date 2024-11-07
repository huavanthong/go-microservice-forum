package services

import (
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/api/models"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/domain/entities"
)

type CategoryService interface {
	CreateCategory(pr *models.RequestCreateCategory) (*entities.Category, error)
	FindAllCategories(page int, limit int) ([]*entities.Category, error)
	FindCategoryByID(id string) (*entities.Category, error)
	FindCategoryByName(name string) ([]*entities.Category, error)
	UpdateCategory(id string, pr *models.RequestUpdateCategory) (*entities.Category, error)
	DeleteCategory(id string) error
}
