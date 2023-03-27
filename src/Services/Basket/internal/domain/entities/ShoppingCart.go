package entities

type ShoppingCart struct {
	UserName string
	Items    []ShoppingCartItem
}

type ShoppingCartItem struct {
	ProductId int
	Price     float64
	Quantity  int
}

func (cart *ShoppingCart) TotalPrice() float64 {
	var totalPrice float64
	for _, item := range cart.Items {
		totalPrice += item.Price * float64(item.Quantity)
	}
	return totalPrice
}
