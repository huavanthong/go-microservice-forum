package repositories

import (
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/domain/entities"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/infrastructure/storage/mongodb"
)

type ProductRepositoryImpl struct {
	productStorage *mongodb.ProductStorage
}

func NewProductRepositoryImpl(productStorage *mongodb.ProductStorage) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{productStorage}
}

func (pr *ProductRepositoryImpl) CreateProduct(p *entities.Product) (*entities.Product, error) {
	return pr.productStorage.Create(p)
}

func (pr *ProductRepositoryImpl) GetProductByID(id string) (*entities.Product, error) {
	return pr.productStorage.GetByID(id)
}

func (pr *ProductRepositoryImpl) UpdateProduct(p *entities.Product) (*entities.Product, error) {
	return pr.productStorage.Update(p)
}

func (pr *ProductRepositoryImpl) DeleteProduct(id string) error {
	return pr.productStorage.Delete(id)
}

func (pr *ProductRepositoryImpl) GetProducts(filter *entities.ProductFilter, pagination *entities.Pagination) ([]*entities.Product, error) {
	return pr.productStorage.GetProducts(filter, pagination)
}
