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

	// Bởi vì chúng ta sẽ không sử dụng việc tính toán totalPrices, totalDiscount ngay lúc này, nên ta sẽ bỏ qua bước này.
	/*
		totalPrice := 0.0
		totalDiscounts := 0.0

		for _, item := range basket.Items {
			totalPrice += item.Price
			totalDiscounts += item.Discount
		}

		basket.TotalPrice = totalPrice
		basket.TotalDiscounts = totalDiscounts
	*/

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

func (bs *BasketServiceImpl) UpdateBasket(ubq *models.UpdateBasketRequest) (*entities.Basket, error) {

	var basket entities.Basket

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
