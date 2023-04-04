package repositories

import (
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/domain/entities"
)

type CategoryRepository interface {
	CreateCategory(c *entities.Category) (*entities.Category, error)
	GetCategoryByID(id string) (*entities.Category, error)
	UpdateCategory(c *entities.Category) (*entities.Category, error)
	DeleteCategory(id string) error
	GetCategories(filter *entities.CategoryFilter, pagination *entities.Pagination) ([]*entities.Category, error)
}

type CategorySearchRepository interface {
	FindCategoryByName(name string) ([]*entities.Category, error)
}
