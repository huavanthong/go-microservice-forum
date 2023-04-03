package storage

import (
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/domain/entities"
)

// ProductStorage is an interface for managing product based as CRUD operation
type ProductStorage interface {
	Create(p *entities.Product) (*entities.Product, error)
	GetByID(id string) (*entities.Product, error)
	Update(p *entities.Product) (*entities.Product, error)
	Delete(id string) error
	GetProducts(filter *entities.ProductFilter, pagination *entities.Pagination) ([]*entities.Product, error)
}
