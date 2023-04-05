package entities

import "errors"

type ShoppingCart struct {
	ShoppingCartID string
	UserName       string
	Items          []ShoppingCartItem
	LoggedIn       bool
}

func NewShoppingCart(id string, name string) *ShoppingCart {
	return &ShoppingCart{
		ShoppingCartID: id,
		UserName:       name,
		Items:          []ShoppingCartItem{},
	}
}

func (cart *ShoppingCart) AddItem(item ShoppingCartItem) error {

	if !cart.LoggedIn {
		return errors.New("User must be logged in to add item to cart")
	}
	cart.Items = append(cart.Items, item)

	return nil
}

func (cart *ShoppingCart) TotalPrice() float64 {
	var totalPrice float64
	for _, item := range cart.Items {
		totalPrice += item.Price * float64(item.Quantity)
	}
	return totalPrice
}
