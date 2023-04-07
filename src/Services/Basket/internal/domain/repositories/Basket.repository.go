package repositories

import (
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/entities"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/interfaces/api/models"
)

type BasketRepository interface {
	CreateBasket(cbr *models.CreateBasketRequest) (*entities.Basket, error)
	GetBasket(userId string) (*entities.Basket, error)
	UpdateBasket(basket *entities.Basket) (*entities.Basket, error)
	DeleteBasket(userName string) error
}
