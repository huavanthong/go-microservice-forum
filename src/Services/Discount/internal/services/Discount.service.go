package services

import (
	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/models"
)

type DiscountService interface {
	GetDiscount(id string) (*models.GetDiscountResponse, error)
	CreateDiscount(discount *models.CreateDiscountRequest) (*models.Discount, error)
	UpdateDiscount(discount *models.UpdateDiscountRequest) (*models.Discount, error)
	DeleteDiscount(id string) error
}
