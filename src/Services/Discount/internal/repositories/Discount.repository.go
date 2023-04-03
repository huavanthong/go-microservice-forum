package repositories

import (
	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/models"
)

type DiscountRepository interface {
	GetDiscountByID(ID int) (*models.DiscountDBResponse, error)
	CreateDiscount(discount *models.Discount) error
	UpdateDiscount(discount *models.Discount) error
	DeleteDiscountByID(ID int) error
}
