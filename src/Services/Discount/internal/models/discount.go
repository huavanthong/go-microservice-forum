package models

type Discount struct {
	ID          int     `json:"id"`
	ProductName string  `json:"product_name"`
	Percentage  float32 `json:"percentage"`
	Quantity    int     `json:"quantity"`
	StartDate   string  `json:"start_date"`
	EndDate     string  `json:"end_date"`
}

// GetDiscountRequest defines the request for getting a discount
type GetDiscountRequest struct {
	ID          string  `json:"id" validate:"required"`
	ProductName string  `json:"product_name" validate:"required"`
	Amount      float32 `json:"amount" validate:"required"`
}

// GetDiscountResponse represents the response from a discount request
type GetDiscountResponse struct {
	ID         string  `json:"id"`
	Percentage float32 `json:"percentage"`
	Quantity   int     `json:"quantity"`
	StartDate  string  `json:"start_date"`
	EndDate    string  `json:"end_date"`
	Available  bool    `json:"available"`
}

// DiscountRequest defines the request for creating or updating a discount
type DiscountRequest struct {
	ProductName string  `json:"product_name" validate:"required"`
	Amount      float32 `json:"amount" validate:"required"`
	StartDate   string  `json:"start_date" validate:"required"`
	EndDate     string  `json:"end_date" validate:"required"`
	Quantity    int     `json:"quantity" validate:"omitempty,gte=1"`
}

// DiscountResponse represents the response from a discount request
type DiscountResponse struct {
	ID          uint    `json:"id"`
	ProductID   uint    `json:"product_id"`
	Percentage  float32 `json:"percentage"`
	Description string  `json:"description"`
}

type GenericResponse struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Errors  []string    `json:"errors,omitempty"`
}
