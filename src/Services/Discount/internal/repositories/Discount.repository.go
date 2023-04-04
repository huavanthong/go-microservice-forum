package repositories

import (
	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/models"
)

type DiscountRepository interface {
	GetDiscount(ID int) (*models.DiscountDBResponse, error)
	CreateDiscount(discount *models.Discount) error
	UpdateDiscount(discount *models.Discount) error
	DeleteDiscount(ID int) error
}
