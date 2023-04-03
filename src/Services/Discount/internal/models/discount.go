package models

type Discount struct {
	ID          int     `json:"id"`
	ProductName string  `json:"product_name"`
	Description string  `json:"description"`
	Percentage  float32 `json:"percentage"`
	Quantity    int     `json:"quantity"`
	StartDate   string  `json:"start_date"`
	EndDate     string  `json:"end_date"`
}

type DiscountDBResponse struct {
	ID          int     `json:"id"`
	ProductName string  `json:"product_name"`
	Description string  `json:"description"`
	Percentage  float32 `json:"percentage"`
	Quantity    int     `json:"quantity"`
	StartDate   string  `json:"start_date"`
	EndDate     string  `json:"end_date"`
	UpdatedAt   string  `json:"updated_at"`
	DeleteAt    string  `json:"deleted_at"`
}

// GetDiscountRequest defines the request for getting a discount
type GetDiscountRequest struct {
	ID          int     `json:"id" validate:"required"`
	ProductName string  `json:"product_name" validate:"required"`
	Amount      float32 `json:"amount" validate:"required"`
}

// GetDiscountResponse represents the response from a discount request
type GetDiscountResponse struct {
	ID          int     `json:"id"`
	ProductName string  `json:"product_name"`
	Description string  `json:"description"`
	Percentage  float32 `json:"percentage"`
	Quantity    int     `json:"quantity"`
	StartDate   string  `json:"start_date"`
	EndDate     string  `json:"end_date"`
	Available   bool    `json:"available"`
}

func FilteredGetResponse(discount *DiscountDBResponse, available bool) *GetDiscountResponse {
	return &GetDiscountResponse{
		ID:          discount.ID,
		ProductName: discount.ProductName,
		Description: discount.Description,
		Percentage:  discount.Percentage,
		Quantity:    discount.Quantity,
		StartDate:   discount.StartDate,
		EndDate:     discount.EndDate,
		Available:   available,
	}
}

// DiscountRequest defines the request for creating or updating a discount
type CreateDiscountRequest struct {
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
