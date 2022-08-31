package models

type Product_thoitrang struct {
	Product
	origin string
}

func NewProductThoiTrang() iProduct {
	return &Product_thoitrang{
		Product: Product{
			Name: "Thoi trang gun",
		},
	}
}
