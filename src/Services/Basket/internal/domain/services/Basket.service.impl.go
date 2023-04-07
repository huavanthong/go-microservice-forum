package services

import (
	"time"

	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/ValueObjects"
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

func convertRequestUpdateToBasket(request *models.UpdateBasketRequest) *entities.Basket {

	basketItems := make([]entities.BasketItem, len(request.Items))
	for i, item := range request.Items {
		basketItems[i] = entities.BasketItem{
			ProductID:      item.ProductID,
			ProductName:    item.ProductName,
			Quantity:       item.Quantity,
			Price:          item.Price,
			DiscountAmount: 0,
			TotalPrice:     float64(item.Quantity) * item.Price,
			ImageURL:       item.ImageURL,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}
	}

	basket := &entities.Basket{
		ID:             ValueObjects.BasketID(request.BasketID),
		UserID:         request.UserID,
		UserName:       request.UserName,
		Items:          basketItems,
		TotalPrice:     0,
		TotalDiscounts: 0,
		UpdatedAt:      time.Now(),
	}

	return basket
}

func (bs *BasketServiceImpl) UpdateBasket(request *models.UpdateBasketRequest) (*entities.Basket, error) {

	// Convert request to update data
	basket := convertRequestUpdateToBasket(request)

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
