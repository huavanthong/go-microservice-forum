package repositories

import (
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/domain/entities"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/infrastructure/storage/mongodb"
)

type CategoryRepositoryImpl struct {
	categoryStorage *mongodb.CategoryStorage
}

func NewCategoryRepositoryImpl(categoryStorage *mongodb.CategoryStorage) CategoryRepository {
	return &CategoryRepositoryImpl{categoryStorage}
}

func (cr *CategoryRepositoryImpl) CreateCategory(c *entities.Category) (*entities.Category, error) {
	return cr.categoryStorage.Create(c)
}

func (cr *CategoryRepositoryImpl) GetCategoryByID(id string) (*entities.Category, error) {
	return cr.categoryStorage.GetByID(id)
}

func (cr *CategoryRepositoryImpl) UpdateCategory(c *entities.Category) (*entities.Category, error) {
	return cr.categoryStorage.Update(c)
}

func (cr *CategoryRepositoryImpl) DeleteCategory(id string) error {
	return cr.categoryStorage.Delete(id)
}

func (cr *CategoryRepositoryImpl) GetCategories(filter *entities.CategoryFilter, pagination *entities.Pagination) ([]*entities.Category, error) {
	return cr.categoryStorage.GetCategories(filter, pagination)
}
