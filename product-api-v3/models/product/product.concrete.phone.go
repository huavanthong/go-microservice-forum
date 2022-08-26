package models

type product_phone struct {
	Product
}

func NewProductPhone() iProduct {
	return &product_phone{
		Product: Product{},
	}
}
