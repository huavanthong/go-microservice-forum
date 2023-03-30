package repositories

import (
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/domain/entities"
)

type ProductRepository interface {
	CreateProduct(p *entities.Product) (*entities.Product, error)
	GetProductByID(id string) (*entities.Product, error)
	UpdateProduct(p *entities.Product) error
	DeleteProduct(id string) error
	GetProducts(filter *entities.ProductFilter, pagination *entities.Pagination) ([]*entities.Product, int64, error)
}

type ProductSearchRepository interface {
	FindProductByName(name string, currency string) ([]*entities.Product, error)
	FindProductByCategory(category string, currency string) ([]*entities.Product, error)
}
