package services

import (
	"context"

	"go.uber.org/zap"

	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/api/models"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/domain/entities"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/infrastruture/storage/mongodb"
)

type CatalogServiceImpl struct {
	log         *zap.Logger
	productRepo *mongodb.ProductRepository
	ctx         context.Context
}

func NewCatalogServiceImpl(log *zap.Logger, productRepo *mongodb.ProductRepository, ctx context.Context) CatalogService {
	return &CatalogServiceImpl{log, productRepo, ctx}
}

func (p *CatalogServiceImpl) CreateProduct(pr *models.RequestCreateProduct) (*entities.Product, error) {

	panic(nil)

}

func (p *CatalogServiceImpl) FindAllProducts(page int, limit int, currency string) (interface{}, error) {

	panic(nil)
}

func (p *CatalogServiceImpl) FindProductByID(id string, currency string) (*entities.Product, error) {
	panic(nil)
}
func (p *CatalogServiceImpl) FindProductByName(name string, currency string) ([]*entities.Product, error) {
	panic(nil)
}

func (p *CatalogServiceImpl) FindProductByCategory(category string, currency string) ([]*entities.Product, error) {

	panic(nil)
}

func (p *CatalogServiceImpl) UpdateProduct(id string, pr *models.RequestUpdateProduct) (*entities.Product, error) {

	panic(nil)
}

func (p *CatalogServiceImpl) DeleteProduct(id string) error {

	panic(nil)
}
