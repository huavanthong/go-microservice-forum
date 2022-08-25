package models

type product_thoitrang struct {
	Product
}

func NewProductThoiTrang() iProduct {
	return &product_thoitrang{
		Product: Product{
			Name: "Thoi trang gun",
		},
	}
}
