package models

type Coupon struct {
	ID          int     `json:"id"`
	ProductName string  `json:"product_name"`
	Description string  `json:"description"`
	Amount      float32 `json:"amount"`
}

// CouponRequest defines the request for creating or updating a coupon
type CouponRequest struct {
	Code       string  `json:"code" validate:"required"`
	Amount     float32 `json:"amount" validate:"required"`
	StartDate  string  `json:"start_date" validate:"required"`
	EndDate    string  `json:"end_date" validate:"required"`
	MinSpend   float32 `json:"min_spend" validate:"omitempty,gte=0"`
	MaxSpend   float32 `json:"max_spend" validate:"omitempty,gte=0"`
	Limit      int     `json:"limit" validate:"omitempty,gte=1"`
	IsPercent  bool    `json:"is_percent"`
	Percentage int     `json:"percentage,omitempty" validate:"omitempty,gte=0,lte=100"`
}

// CouponResponse represents the response from a coupon request
type CouponResponse struct {
	ID          uint    `json:"id"`
	Code        string  `json:"code"`
	Percentage  float32 `json:"percentage"`
	Description string  `json:"description"`
	MaxUses     int     `json:"max_uses"`
	ExpiresAt   int64   `json:"expires_at"`
}
