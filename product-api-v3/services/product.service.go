package services

import "github.com/huavanthong/microservice-golang/product-api-v3/models"

type ProductService interface {
	FindAllProducts(page int, limit int, currency string) (models.Products, error)
	FindProductByID(id string, currency string) (*models.Product, error)
	FindProductByName(name string, currency string) (*models.Product, error)
	FindProductByCategory(category string, currency string) (*models.Product, error)
	CreateProduct(pr *models.Product) (*models.Product, error)
	UpdateProduct(pr models.Product) (*models.Product, error)
	DeleteProduct(id string) error
}
