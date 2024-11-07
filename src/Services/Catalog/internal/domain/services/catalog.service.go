package services

import (
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/api/models"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/domain/entities"
)

// CatalogService defines the interface for catalog service.
type CatalogService interface {
	// CatalogRepository
	CreateProduct(pr *models.CreateProductRequest) (*entities.Product, error)
	GetProductByID(id string, currency string) (*entities.Product, error)
	UpdateProduct(id string, pr *models.UpdateProductRequest) (*entities.Product, error)
	DeleteProduct(id string) error
	GetProducts(page int, limit int, currency string) (interface{}, error)

	// CatalogSearchRepository
	FindProductByName(name string, currency string) ([]*entities.Product, error)
	FindProductByCategory(category string, currency string) ([]*entities.Product, error)
}
