package models

import "fmt"

func getProductType(productType string) (iProduct, error) {
	if productType == "phone" {
		return NewProductPhone(), nil
	}
	if productType == "dien-tu" {
		return NewProductDienTu(), nil
	}
	if productType == "cloths" {
		return NewProductThoiTrang(), nil
	}
	return nil, fmt.Errorf("Wrong product type passed")
}
