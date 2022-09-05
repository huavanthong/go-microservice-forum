package models

type Product_thoitrang struct {
	Product
	Material string
}

func NewProductThoiTrang() iProduct {
	return &Product_thoitrang{
		Product:  Product{},
		Material: "Cotton",
	}
}
