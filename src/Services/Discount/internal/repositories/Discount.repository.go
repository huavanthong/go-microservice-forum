package repositories

import (
	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/models"
)

type DiscountRepository interface {
	GetDiscounts() ([]*models.Discount, error)
	GetDiscount(ID int) (*models.Discount, error)
	CreateDiscount(discount *models.Discount) (*models.Discount, error)
	UpdateDiscount(discount *models.Discount) (*models.Discount, error)
	DeleteDiscount(ID int) error
}
