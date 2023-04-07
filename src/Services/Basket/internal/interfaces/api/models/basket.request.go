package models

type CreateBasketRequest struct {
	UserID   string `json:"user_id" validate:"required" example:"1234567890"`
	UserName string `json:"user_name" validate:"required" example:"hvthong"`
}

type UpdateBasketRequest struct {
	UserID   string `json:"user_id" validate:"required" example:"1234567890"`
	UserName string `json:"user_name" validate:"required" example:"hvthong"`
	Items    []Item `json:"items"`
}
type Item struct {
	ProductID   string  `json:"product_id" validate:"required" `
	ProductName string  `json:"product_name" validate:"required" example:"phone"`
	Quantity    int     `json:"quantity" validate:"required,min=1" example:"1"`
	Price       float64 `json:"price" example:"1400" `
	ImageURL    string  `json:"image_url" example:"default.png"`
}
