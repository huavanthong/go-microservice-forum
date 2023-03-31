package models

type Coupon struct {
	ID          int     `json:"id"`
	ProductName string  `json:"product_name"`
	Description string  `json:"description"`
	Amount      float32 `json:"amount"`
}
