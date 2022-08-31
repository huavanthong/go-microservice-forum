package models

type Product_dientu struct {
	Product
}

func NewProductDienTu() iProduct {
	return &Product_dientu{
		Product: Product{
			Name: "AK47 gun",
		},
	}
}
