package entities

type ShoppingCart struct {
	UserName string
	Items    []ShoppingCartItem
}

func NewShoppingCart(userName string) *ShoppingCart {
	return &ShoppingCart{
		UserName: userName,
		Items:    []ShoppingCartItem{},
	}
}
func (cart *ShoppingCart) TotalPrice() float64 {
	var totalPrice float64
	for _, item := range cart.Items {
		totalPrice += item.Price * float64(item.Quantity)
	}
	return totalPrice
}
