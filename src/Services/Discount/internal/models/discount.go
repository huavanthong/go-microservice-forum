package models

import (
	"time"

	"github.com/google/uuid"
)

type Discount struct {
	ID           uuid.UUID `gorm:"column:id;type:uuid;default:uuid_generate_v4();primary_key" json:"id, omitempty" `
	ProductID    string    `gorm:"column:product_id;varchar(255);uniqueIndex;not null" json:"product_id"`
	ProductName  string    `gorm:"column:product_name;not null" json:"product_name"`
	Description  string    `gorm:"column:description;varchar(255)" json:"description"`
	DiscountType string    `gorm:"column:discount_type;varchar(100)" json:"discount_type"`
	Percentage   float32   `gorm:"column:percentage" json:"percentage"`
	Amount       float64   `gorm:"column:amount" json:"amount"`
	Quantity     int       `gorm:"column:quantity" json:"quantity"`
	StartDate    string    `gorm:"column:start_date; not null" json:"start_date"`
	EndDate      string    `gorm:"column:end_date; not null" json:"end_date"`
	CreatedAt    time.Time `gorm:"column:created_at; not null" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at; not null" json:"updated_at"`
}

type GenericResponse struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Errors  []string    `json:"errors,omitempty"`
}
