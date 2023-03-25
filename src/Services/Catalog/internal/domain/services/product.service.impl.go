package services

import (
	"context"

	"go.uber.org/zap"

	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/api/models"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/domain/entities"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/domain/repositories"
)

type ProductServiceImpl struct {
	log         *zap.Logger
	productRepo *repositories.ProductRepository
	ctx         context.Context
}

func NewProductServiceImpl(log *zap.Logger, productRepo *repositories.ProductRepository, ctx context.Context) ProductService {
	return &ProductServiceImpl{log, productRepo, ctx}
}

func (p *ProductServiceImpl) CreateProduct(pr *models.RequestCreateProduct) (*entities.Product, error) {

	panic(nil)

}

func (p *ProductServiceImpl) FindAllProducts(page int, limit int, currency string) (interface{}, error) {

	panic(nil)
}

func (p *ProductServiceImpl) FindProductByID(id string, currency string) (*entities.Product, error) {
	panic(nil)
}
func (p *ProductServiceImpl) FindProductByName(name string, currency string) ([]*entities.Product, error) {
	panic(nil)
}

func (p *ProductServiceImpl) FindProductByCategory(category string, currency string) ([]*entities.Product, error) {

	panic(nil)
}

func (p *ProductServiceImpl) UpdateProduct(id string, pr *models.RequestUpdateProduct) (*entities.Product, error) {

	panic(nil)
}

func (p *ProductServiceImpl) DeleteProduct(id string) error {

	panic(nil)
}
