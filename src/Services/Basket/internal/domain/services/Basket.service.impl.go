package services

import (
	"errors"
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

func convertRequestCreateToBasket(request *models.CreateBasketRequest) *entities.Basket {

	basketItems := make([]entities.BasketItem, 0)
	createdAt := time.Now()

	basket := &entities.Basket{
		ID:             ValueObjects.NewBasketID(),
		UserID:         request.UserID,
		UserName:       request.UserName,
		Items:          basketItems,
		TotalPrice:     0,
		TotalDiscounts: 0,
		CreatedAt:      createdAt,
		UpdatedAt:      createdAt,
		ExpiresAt:      createdAt,
	}

	return basket
}

func (bs *BasketServiceImpl) CreateBasket(cbr *models.CreateBasketRequest) (*entities.Basket, error) {

	// Convert request to create basket data
	basketRequest := convertRequestCreateToBasket(cbr)
	basket, err := bs.basketRepo.CreateBasket(basketRequest)
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

func convertRequestUpdateToBasket(request *models.UpdateBasketRequest, oldBasket *entities.Basket) *entities.Basket {
	createdAt := time.Now()

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
			CreatedAt:      createdAt,
			UpdatedAt:      createdAt,
		}

		if oldBasket != nil {
			for _, oldItem := range oldBasket.Items {
				if oldItem.ProductID == item.ProductID {
					basketItems[i].CreatedAt = oldItem.CreatedAt
					break
				}
			}
		}
	}

	basket := &entities.Basket{
		ID:             ValueObjects.BasketID(request.BasketID),
		UserID:         request.UserID,
		UserName:       request.UserName,
		Items:          basketItems,
		TotalPrice:     0,
		TotalDiscounts: 0,
		CreatedAt:      oldBasket.CreatedAt,
		UpdatedAt:      createdAt,
		ExpiresAt:      oldBasket.ExpiresAt,
	}

	return basket
}

func (bs *BasketServiceImpl) UpdateBasket(request *models.UpdateBasketRequest) (*entities.Basket, error) {

	// Check basket exist by user ID
	basket, err := bs.basketRepo.GetBasket(request.UserID)
	if err != nil {
		return nil, err
	}

	if basket == nil {
		return nil, errors.New("Basket not found by user id")
	}

	// Convert request to update basket data
	basketRequest := convertRequestUpdateToBasket(request, basket)

	// Execute to update basket
	basket, err = bs.basketRepo.UpdateBasket(basketRequest)
	if err != nil {
		return nil, err
	}
	return basket, nil
}

func (bs *BasketServiceImpl) DeleteBasket(userName string) error {
	err := bs.basketRepo.DeleteBasket(userName)
	if err != nil {
		return err
	}
	return nil
}
