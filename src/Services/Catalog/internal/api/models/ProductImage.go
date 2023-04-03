package models

import "github.com/google/uuid"

type ProductImage struct {
	ID       uuid.UUID
	ImageURL string
	IsMain   bool
}

func NewProductImage(id uuid.UUID, imageURL string, isMain bool, productID uuid.UUID) *ProductImage {
	return &ProductImage{
		ID:       id,
		ImageURL: imageURL,
		IsMain:   isMain,
	}
}

func (pi *ProductImage) SetIsMain(isMain bool) {
	pi.IsMain = isMain
}

func (pi *ProductImage) SetImageURL(url string) {
	pi.ImageURL = url
}

type EntityId uuid.UUID
