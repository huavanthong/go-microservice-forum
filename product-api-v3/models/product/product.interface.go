package models

import (
	"github.com/huavanthong/microservice-golang/product-api-v3/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type iProduct interface {
	GetID() primitive.ObjectID
	SetID(ID primitive.ObjectID)

	GetProductCode() string
	SetProductCode(ProductCode string)

	GetName() string
	SetName(Name string)

	GetCategory() models.Category
	SetCategory(Category models.Category)

	GetInventory() models.Inventory
	SetInventory(Inventory models.Inventory)

	GetBrand() models.Brand
	SetBrand(Brand models.Brand)

	GetSummary() string
	SetSummary(Summary string)

	GetDescription() string
	SetDescription(Description string)

	GetImageFile() string
	SetImageFile(ImageFile string)

	GetDiscount() models.Discount
	SetDiscount(Discount models.Discount)

	GetPrice() float64
	SetPrice(Price float64)

	GetSKU() string
	SetSKU(SKU string)

	GetCreatedAt() string
	SetCreatedAt(CreatedAt string)

	GetUpdatedAt() string
	SetUpdatedAt(UpdatedAt string)

	GetDeleteAt() string
	SetDeleteAt(DeleteAt string)
}
