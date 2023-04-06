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

	basket, err := bs.basketRepo.CreateBasket(cbr)
	if err != nil {
		return nil, err
	}
	return basket, nil
}

func (bs *BasketServiceImpl) GetBasket(userId string) (*entities.Basket, error) {
	basket, err := bs.basketRepo.GetBasket(userId)
	if err != nil {
		return nil, err
	}
	return basket, nil
}

func (bs *BasketServiceImpl) UpdateBasket(basket *entities.Basket) (*entities.Basket, error) {
	updatedBasket, err := bs.basketRepo.UpdateBasket(basket)
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
