package models

type product_thoitrang struct {
	product
}


func NewProductThoiTrang iProduct {
	return &product_thoitrang{
        product: product{
            Name:  "Thoi trang gun",
        },
    }
}