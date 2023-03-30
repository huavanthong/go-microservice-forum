package repositories

import (
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/domain/entities"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/infrastructure/storage"
)

type ProductRepositoryImpl struct {
	productStorage *storage.ProductStorage
}

func NewProductRepositoryImpl(productStorage *storage.ProductStorage) ProductRepository {
	return &ProductRepositoryImpl{productStorage}
}

func (pr *ProductRepositoryImpl) CreateProduct(p *entities.Product) (*entities.Product, error) {
	return p.repository.Create(p)
}

func (pr *ProductRepositoryImpl) GetProductByID(id string) (*entities.Product, error) {
	return pr.productStorage.GetByID(id)
}

func (pr *ProductRepositoryImpl) UpdateProduct(p *entities.Product) error {
	return p.productStorage.Update(p)
}

func (pr *ProductRepositoryImpl) DeleteProduct(id string) error {
	return pr.productStorage.Delete(id)
}

func (pr *ProductRepositoryImpl) GetProducts(filter *entities.ProductFilter, pagination *entities.Pagination) ([]*entities.Product, int64, error) {
	return pr.productStorage.GetByID(filter, pagination)
}
