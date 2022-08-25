package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/************************ Define structure product ************************/
type product struct {
	ID          primitive.ObjectID `bson:"_id" json:"id" example:"5bbdadf782ebac06a695a8e7"`
	ProductCode string             `bson:"pcode" json:"pcode" binding:"required" example:"p123456789"`
	Name        string             `json:"name" bson:"name" binding:"required,gt=0,lt=255" example:"Iphone 14 Pro"`
	Category    string             `json:"category" bson:"category" binding:"required,gt=0,lt=255" example:"Phone"`
	Summary     string             `json:"summary" bson:"summary" binding:"required,gt=0,lt=10000" example:"Iphone 14 Pro Gold"`
	Description string             `json:"description" bson:"description" binding:"required,gt=0,lt=10000" example:"Iphone 14 Pro Gold 256GB"`
	ImageFile   string             `json:"imageFile" bson:"imageFile" binding:"required" example:"default.png"`
	Price       float64            `json:"price" bson:"price" binding:"required,min=0.01" example:"1400"`
	SKU         string             `json:"sku" bson:"sku" example:"ABC-XYZ-OXY"`
	CreatedAt   string             `json:"created_at" bson:"created_at"`
	UpdatedAt   string             `json:"updated_at" bson:"updated_at"`
	DeleteAt    string             `json:"deleted_at" bson:"deleted_at"`
}

/************************ Implement setter/getter method for Product ************************/
func (product *product) GetID() primitive.ObjectID {
	return product.ID
}

func (product *product) SetID(ID primitive.ObjectID) *product {
	product.ID = ID
	return product
}

func (product *product) GetProductCode() string {
	return product.ProductCode
}

func (product *product) SetProductCode(ProductCode string) *product {
	product.ProductCode = ProductCode
	return product
}

func (product *product) GetName() string {
	return product.Name
}

func (product *product) SetName(Name string) *product {
	product.Name = Name
	return product
}

func (product *product) GetCategory() string {
	return product.Category
}

func (product *product) SetCategory(Category string) *product {
	product.Category = Category
	return product
}

func (product *product) GetSummary() string {
	return product.Summary
}

func (product *product) SetSummary(Summary string) *product {
	product.Summary = Summary
	return product
}

func (product *product) GetDescription() string {
	return product.Description
}

func (product *product) SetDescription(Description string) *product {
	product.Description = Description
	return product
}

func (product *product) GetImageFile() string {
	return product.ImageFile
}

func (product *product) SetImageFile(ImageFile string) *product {
	product.ImageFile = ImageFile
	return product
}

func (product *product) GetPrice() float64 {
	return product.Price
}
func (product *product) SetPrice(Price float64) *product {
	product.Price = Price
	return product
}

func (product *product) GetSKU() string {
	return product.SKU
}
func (product *product) SetSKU(SKU string) *product {
	product.SKU = SKU
	return product
}

func (product *product) GetCreatedAt() string {
	return product.CreatedAt
}

func (product *product) SetCreatedAt(CreatedAt string) *product {
	product.CreatedAt = CreatedAt
	return product
}
func (product *product) GetUpdatedAt() string {
	return product.UpdatedAt
}

func (product *product) SetUpdatedAt(UpdatedAt string) *product {
	product.UpdatedAt = UpdatedAt
	return product
}

func (product *product) GetDeleteAt() string {
	return product.DeleteAt
}

func (product *product) SetDeleteAt(DeleteAt string) *product {
	product.DeleteAt = DeleteAt
	return product
}

/************************ Filter info ************************/
func FilteredResponse(p *product) product {
	return product{
		ID:          p.ID,
		ProductCode: p.ProductCode,
		Name:        p.Name,
		Category:    p.Category,
		Summary:     p.Summary,
		Description: p.Description,
		ImageFile:   p.ImageFile,
		Price:       p.Price,
		SKU:         p.SKU,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}
