package models

type Product_phone struct {
	Product
	Model string
}

func NewProductPhone() iProduct {
	return &Product_phone{
		Product: Product{},
		Model:   "test",
	}
}
