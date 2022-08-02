package services

import (
	"context"

	"github.com/huavanthong/microservice-golang/product-api-v3/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewProductServiceImpl(collection *mongo.Collection, ctx context.Context) UserService {
	return &ProductServiceImpl{collection, ctx}
}

func (ps *ProductServiceImpl) FindAllProducts(page int, limit int, currency string) (models.Products, error) {
	panic("Implement me")
}

func (ps *ProductServiceImpl) FindProductByID(id string, currency string) (*models.Product, error) {

	panic("Implement me")
}
func (ps *ProductServiceImpl) FindProductByName(name string, currency string) (*models.Product, error) {

	panic("Implement me")
}
func (ps *ProductServiceImpl) FindProductByCategory(name string, currency string) (*models.Product, error) {

	panic("Implement me")
}
func (ps *ProductServiceImpl) CreateProduct(pr *models.Product) error {

	panic("Implement me")
}
func (ps *ProductServiceImpl) UpdateProduct(pr models.Product) error {

	panic("Implement me")
}
func (ps *ProductServiceImpl) DeleteProduct(id string) error {

	panic("Implement me")
}
