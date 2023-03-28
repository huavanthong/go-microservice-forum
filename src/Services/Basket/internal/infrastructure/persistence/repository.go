package persistence

import (
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/entities"
)

type BasketRepository interface {
	GetBasket(userName string) (*entities.ShoppingCart, error)
	UpdateBasket(basket *entities.ShoppingCart) (*entities.ShoppingCart, error)
	DeleteBasket(userName string) error
}
