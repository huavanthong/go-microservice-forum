package models

type product_phone struct {
	Product
	Model string
}

func NewProductPhone() iProduct {
	return &product_phone{
		Product: Product{},
		Model:   "test",
	}
}
