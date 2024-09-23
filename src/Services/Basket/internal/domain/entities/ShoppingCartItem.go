package entities

type ShoppingCartItem struct {
	Quantity    int     `json:"quantity"`
	Color       string  `json:"color"`
	Price       float64 `json:"price"`
	ProductID   string  `json:"productId"`
	ProductName string  `json:"productName"`
}
