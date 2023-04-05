package repositories

import (
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/entities"
)

type BasketRepository interface {
	CreateBasket(userId string) (*entities.Basket, error)
	GetBasket(userName string) (*entities.Basket, error)
	UpdateBasket(cart *entities.Basket) (*entities.Basket, error)
	DeleteBasket(userName string) error
}
