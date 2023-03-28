package repositories

import (
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/entities"
)

// BasketRepository is an interface for managing Basket
type BasketRepository interface {
	Create(userName string) (*entities.ShoppingCart, error)
	GetByUserName(userName string) (*entities.ShoppingCart, error)
	Update(basket *entities.ShoppingCart) (*entities.ShoppingCart, error)
	Delete(userName string) error
}
