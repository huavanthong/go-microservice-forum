package repositories

import (
	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/models"
)

type DiscountRepository interface {
	GetDiscount(productName string) models.Coupon
	CreateDiscount(coupon models.Coupon) bool
	UpdateDiscount(coupon models.Coupon) bool
	DeleteDiscount(productName string) bool
}
