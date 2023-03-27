package services

import (
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/entities"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/repositories"
)

type BasketService struct {
	repo repositories.BasketRepository
}

func NewBasketService(repo repositories.BasketRepository) *BasketService {
	return &BasketService{repo: repo}
}

func (s *BasketService) GetBasket(userName string) (*entities.ShoppingCart, error) {
	return s.repo.GetBasket(userName)
}

func (s *BasketService) UpdateBasket(basket *entities.ShoppingCart) error {
	return s.repo.UpdateBasket(basket)
}

func (s *BasketService) DeleteBasket(userName string) error {
	return s.repo.DeleteBasket(userName)
}
