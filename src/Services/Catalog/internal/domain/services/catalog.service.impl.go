package services

import (
	"context"
	"errors"

	"go.uber.org/zap"

	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/api/models"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/domain/builders"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/domain/entities"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/domain/repository"
)

type CatalogServiceImpl struct {
	log         *zap.Logger
	catalogRepo *repository.ProductRepository
	ctx         context.Context
}

func NewCatalogServiceImpl(log *zap.Logger, productRepo *repository.ProductRepository, ctx context.Context) CatalogService {
	return &CatalogServiceImpl{log, productRepo, ctx}
}

func (cs *CatalogServiceImpl) CreateProduct(pr *models.CreateProductRequest) (*entities.Product, error) {

	// Initialize product builder
	productBuilder := builders.NewProductBuilder()

	// Mapping request info to product builder
	productBuilder.SetName(pr.Name)
	productBuilder.SetPrice(pr.Price)
	productBuilder.SetCategory(pr.Category)
	productBuilder.SetBrand(pr.Brand)
	productBuilder.SetSummary(pr.Summary)
	productBuilder.SetDescription(pr.Description)
	productBuilder.SetImageFile(pr.ImageFile)
	product := productBuilder.Build()

	productRes, err := cs.catalogRepo.CreateProduct(product)
	if err != nil {
		return nil, errors.New("Failed to create new product")
	}

	return productRes, nil
}

func (p *CatalogServiceImpl) GetProducts(page int, limit int, currency string) (interface{}, error) {

	panic(nil)
}

func (cs *CatalogServiceImpl) GetProductByID(id string, currency string) (*entities.Product, error) {

	return cs.catalogRepo.GetProductByID(id)
}

func (cs *CatalogServiceImpl) UpdateProduct(id string, pr *models.RequestUpdateProduct) (*entities.Product, error) {

	// Initialize product builder
	productBuilder := builders.NewProductBuilder()

	// Mapping request info to product builder
	productBuilder.SetName(pr.Name)
	productBuilder.SetPrice(pr.Price)
	productBuilder.SetCategory(pr.Category)
	productBuilder.SetBrand(pr.Brand)
	productBuilder.SetSummary(pr.Summary)
	productBuilder.SetDescription(pr.Description)
	productBuilder.SetImageFile(pr.ImageFile)
	product := productBuilder.Build()

	productRes, err := cs.catalogRepo.UpdateProduct(product)
	if err != nil {
		return nil, errors.New("Failed to update new product")
	}

	return productRes, nil
}

func (cs *CatalogServiceImpl) DeleteProduct(id string) error {

	return cs.catalogRepo.GetProductByID(id)
}
