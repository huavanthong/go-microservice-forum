package services

import (
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/entities"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/repositories"

	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/interfaces/api/models"
)

type BasketServiceImpl struct {
	basketRepo repositories.BasketRepository
}

func NewBasketServiceImpl(basketRepo repositories.BasketRepository) BasketService {
	return &BasketServiceImpl{
		basketRepo: basketRepo,
	}
}

func (bs *BasketServiceImpl) CreateBasket(cbr *models.CreateBasketRequest) (*entities.Basket, error) {

	basket, err := bs.basketRepo.CreateBasket(cbr.UserID)
	if err != nil {
		return nil, err
	}

	// Generate response
	basket.UserName = cbr.UserName

	return basket, nil
}

func (bs *BasketServiceImpl) GetBasket(userName string) (*entities.Basket, error) {
	basket, err := bs.basketRepo.GetBasket(userName)
	if err != nil {
		return nil, err
	}
	return basket, nil
}

func (bs *BasketServiceImpl) UpdateBasket(userName string, cart *entities.Basket) (*entities.Basket, error) {
	updatedBasket, err := bs.basketRepo.UpdateBasket(userName, cart)
	if err != nil {
		return nil, err
	}
	return updatedBasket, nil
}

func (bs *BasketServiceImpl) DeleteBasket(userName string) error {
	err := bs.basketRepo.DeleteBasket(userName)
	if err != nil {
		return err
	}
	return nil
}
