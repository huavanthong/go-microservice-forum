package models

// UpdateDiscountRequest defines the request for updating a discount
type UpdateDiscountRequest struct {
	ID           string  `json:"id" validate:"required" example:"1"`
	ProductID    string  `json:"product_id" validate:"required" example:"5bbdadf782ebac06a695a8e7"`
	ProductName  string  `json:"product_name" example:"laptopn thinkpad"`
	Description  string  `json:"description" example:"black friday"`
	DiscountType string  `json:"discount_type" validate:"required" example:"percent | amount"`
	Percentage   float32 `json:"percentage" example:"10"`
	Amount       float64 `json:"amount" example:"15"`
	Quantity     int     `json:"quantity" validate:"required" example:"100"`
	StartDate    string  `json:"start_date" validate:"required" example:"13-04-2023"`
	EndDate      string  `json:"end_date" validate:"required" example:"25-04-2023"`
}

// UpdateDiscountResponse represents the response from a discount request
type UpdateDiscountResponse struct {
	ID          int     `json:"id"`
	ProductID   string  `json:"product_id"`
	ProductName string  `json:"product_name"`
	Description string  `json:"description"`
	Percentage  float32 `json:"percentage"`
	Amount      float64 `json:"amount"`
	Quantity    int     `json:"quantity"`
	StartDate   string  `json:"start_date"`
	EndDate     string  `json:"end_date"`
}
