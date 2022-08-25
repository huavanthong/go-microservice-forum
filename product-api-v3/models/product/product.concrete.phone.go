package models

type product_phone struct {
	product
}

func NewProductPhone iProduct {
	return &product_phone{
        product: product{
            Name:  "Phone",
        },
    }
}