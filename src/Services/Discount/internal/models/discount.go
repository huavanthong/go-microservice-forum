package models

import "time"

type Discount struct {
	ID           int       `json:"id" db:"id"`
	ProductID    string    `json:"product_id" db:"product_id"`
	ProductName  string    `json:"product_name" db:"product_name"`
	Description  string    `json:"description" db:"description"`
	DiscountType string    `json:"discount_type" db:"discount_type"`
	Percentage   float32   `json:"percentage" db:"percentage"`
	Amount       float64   `json:"amount" db:"amount"`
	Quantity     int       `json:"quantity" db:"quantity"`
	StartDate    string    `json:"start_date" db:"start_date"`
	EndDate      string    `json:"end_date" db:"end_date"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

type GenericResponse struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Errors  []string    `json:"errors,omitempty"`
}
