package repositories

import (
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/entities"
)

type BasketRepository interface {
	CreateBasket(basket *entities.Basket) (*entities.Basket, error)
	GetBasket(userId string) (*entities.Basket, error)
	UpdateBasket(basket *entities.Basket) (*entities.Basket, error)
	DeleteBasket(userName string) error
}
