package models

type CreateBasketResponse struct {
	UserID   string         `json:"user_id"`
	UserName string         `json:"user_name"`
	Items    []ItemResponse `json:"items"`
}

type ItemResponse struct {
	ProductID   string  `json:"product_id" `
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	ImageURL    string  `json:"image_url"`
}
