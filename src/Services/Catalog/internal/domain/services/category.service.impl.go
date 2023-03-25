package services

import (
	"context"

	"go.uber.org/zap"

	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/api/models"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/domain/entities"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/domain/repositories"
)

type CategoryServiceImpl struct {
	log          *zap.Logger
	categoryRepo repositories.CategoryRepository
	ctx          context.Context
}

func NewCategoryServiceImpl(log *zap.Logger, categoryRepo repositories.CategoryRepository, ctx context.Context) CategoryService {
	return &CategoryServiceImpl{log, categoryRepo, ctx}
}

func (c *CategoryServiceImpl) CreateCategory(rc *models.RequestCreateCategory) (*entities.Category, error) {
	panic(nil)
}

func (c *CategoryServiceImpl) FindAllCategories(page int, limit int) ([]*entities.Category, error) {

	panic(nil)
}

func (c *CategoryServiceImpl) FindCategoryByID(id string) (*entities.Category, error) {
	panic(nil)
}

func (c *CategoryServiceImpl) FindCategoryByName(name string) ([]*entities.Category, error) {

	panic(nil)
}

func (c *CategoryServiceImpl) UpdateCategory(id string, pr *models.RequestUpdateCategory) (*entities.Category, error) {

	panic(nil)
}

func (c *CategoryServiceImpl) DeleteCategory(id string) error {

	panic(nil)
}
