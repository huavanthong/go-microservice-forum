package models

// GetDiscountRequest defines the request for getting a discount
type GetDiscountRequest struct {
	ID        string `json:"id" validate:"required"`
	ProductID string `json:"product_id" validate:"required"`
}

// GetDiscountResponse represents the response from a discount request
type GetDiscountResponse struct {
	ID           string  `json:"id"`
	ProductName  string  `json:"product_name"`
	Description  string  `json:"description"`
	DiscountType string  `json:"discount_type"`
	Percentage   float32 `json:"percentage"`
	Amount       float64 `json:"amount"`
	Quantity     int     `json:"quantity"`
	StartDate    string  `json:"start_date"`
	EndDate      string  `json:"end_date"`
	Available    bool    `json:"available"`
}
