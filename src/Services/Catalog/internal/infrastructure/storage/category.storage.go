package storage

import (
	"Services/Catalog/internal/domain/entities"

	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/domain/entities"
)

// CategoryRepository is an interface for managing product based as CRUD operation
type CategoryStorage interface {
	Create(p *entities.Category) (*entities.Category, error)
	GetByID(id string) (*entities.Category, error)
	Update(p *entities.Category) error
	Delete(id string) error
	GetCategories(filter *entities.CategoryFilter, pagination *entities.Pagination) ([]*entities.Category, int64, error)
}
