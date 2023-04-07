package persistence

import (
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/entities"
)

// BasketPersistence is an interface for managing Basket based on CRUD operation
type BasketPersistence interface {
	Create(basket *entities.Basket) (*entities.Basket, error)
	Get(userId string) (*entities.Basket, error)
	Update(basket *entities.Basket) (*entities.Basket, error)
	Delete(userId string) error
}
