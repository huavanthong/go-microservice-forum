package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type iProduct interface {
	GetID() primitive.ObjectID
	SetID(ID primitive.ObjectID) *product
	GetProductCode() string
	SetProductCode(ProductCode string) *product
	GetName() string
	SetName(Name string) *product
	GetCategory() string
	SetCategory(Category string) *product
	GetSummary() string
	SetSummary(Summary string) *product
	GetDescription() string
	SetDescription(Description string) *product
	GetImageFile() string
	SetImageFile(ImageFile string) *product
	GetPrice() float64
	SetPrice(Price float64) *product
	GetSKU() string
	SetSKU(SKU string) *product
	GetCreatedAt() string
	SetCreatedAt(CreatedAt string) *product
	GetUpdatedAt() string
	SetUpdatedAt(UpdatedAt string) *product
	GetDeleteAt() string
	SetDeleteAt(DeleteAt string) *product
}
