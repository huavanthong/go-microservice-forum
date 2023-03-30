package services

import (
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/api/models"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/domain/entities"
)

type CatalogService interface {
	CreateProduct(pr *models.RequestCreateProduct) (*entities.Product, error)
	FindAllProducts(page int, limit int, currency string) (interface{}, error)
	FindProductByID(id string, currency string) (*entities.Product, error)
	FindProductByName(name string, currency string) ([]*entities.Product, error)
	FindProductByCategory(category string, currency string) ([]*entities.Product, error)
	UpdateProduct(id string, pr *models.RequestUpdateProduct) (*entities.Product, error)
	DeleteProduct(id string) error
}
