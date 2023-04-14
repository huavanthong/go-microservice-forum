package repositories

import (
	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/models"
)

type DiscountRepository interface {
	GetDiscount(id string) (*models.Discount, error)
	CreateDiscount(discount *models.Discount) (*models.Discount, error)
	UpdateDiscount(discount *models.Discount) (*models.Discount, error)
	DeleteDiscount(id string) error
}
