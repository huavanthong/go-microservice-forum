package models

/************************ Define structure product ************************/
type ShoppingCartItem struct {
	Quantity    int     `json:"quantity" bson:"quantity"`
	Color       int     `json:"color" bson:"color"`
	Price       float32 `json:"price" bson:"price"`
	ProductId   string  `json:"productId" bson:"productId"`
	ProductName string  `json:"productName" bson:"productName"`
}
