package models

type product_dientu struct {
	Product
}

func NewProductDienTu() iProduct {
	return &product_dientu{
		Product: Product{
			Name: "AK47 gun",
		},
	}
}
