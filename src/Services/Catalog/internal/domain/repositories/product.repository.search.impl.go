package repositories

import (
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/domain/entities"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/infrastructure/storage/mongodb"
)

type ProductSearchRepositoryImpl struct {
	productStorage *mongodb.ProductStorage
}

func NewProductSearchRepositoryImpl(productStorage *mongodb.ProductStorage) *ProductSearchRepositoryImpl {
	return &ProductSearchRepositoryImpl{productStorage}
}

func (pr *ProductSearchRepositoryImpl) FindProductByName(name string, currency string) ([]*entities.Product, error) {

	// Todo: need implement logic for currency
	return pr.productStorage.GetProducts(nil, nil)
}

func (pr *ProductSearchRepositoryImpl) FindProductByCategory(category string, currency string) ([]*entities.Product, error) {
	//filter := bson.M{"categories": bson.M{"$elemMatch": bson.M{"name": category}}}

	// Todo: need implement logic for currency
	return pr.productStorage.GetProducts(nil, nil)
}
