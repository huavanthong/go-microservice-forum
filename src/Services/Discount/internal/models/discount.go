package models

type Discount struct {
	ID          int     `json:"id"`
	ProductName string  `json:"product_name"`
	Percentage  float32 `json:"percentage"`
	StartDate   string  `json:"start_date"`
	EndDate     string  `json:"end_date"`
}
