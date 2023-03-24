package models

import "fmt"

// Define Product Type
type ProductType string

// Define constant value
const (
	Phone  ProductType = "phone"
	DienTu             = "dien-tu"
	Cloths             = "thoi-trang"
)

func GetProductType(ptype ProductType) (iProduct, error) {

	switch ptype {
	case Phone:
		return NewProductPhone(), nil
	case DienTu:
		return NewProductDienTu(), nil
	case Cloths:
		return NewProductThoiTrang(), nil
	default:
		return nil, fmt.Errorf("Wrong product type passed")
	}
}
