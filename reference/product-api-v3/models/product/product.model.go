package models

import (
	"github.com/huavanthong/microservice-golang/product-api-v3/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/************************ Define structure product ************************/
type Product struct {
	ID          primitive.ObjectID `bson:"_id" json:"id" example:"5bbdadf782ebac06a695a8e7"`
	ProductCode string             `bson:"pcode" json:"pcode" binding:"required" example:"p123456789"`
	ProductType string             `bson:"ptype" json:"ptype" binding:"required" example:"phone"`
	Name        string             `json:"name" bson:"name" binding:"required,gt=0,lt=255" example:"Iphone 14 Pro"`
	Category    models.Category    `json:"category" bson:"category" binding:"required,gt=0,lt=255"`
	Inventory   models.Inventory   `json:"inventory" bson:"inventory" binding:"required"`
	Brand       models.Brand       `json:"brand" bson:"brand" binding:"required"`
	Summary     string             `json:"summary" bson:"summary" binding:"required,gt=0,lt=10000" example:"Iphone 14 Pro Gold"`
	Description string             `json:"description" bson:"description" binding:"required,gt=0,lt=10000" example:"Iphone 14 Pro Gold 256GB"`
	ImageFile   string             `json:"imageFile" bson:"imageFile" binding:"required" example:"default.png"`
	Discount    models.Discount    `json:"discount" bson:"discount" binding:"required"`
	Price       float64            `json:"price" bson:"price" binding:"required,min=0.01" example:"1400"`
	SKU         string             `json:"sku" bson:"sku" example:"ABC-XYZ-OXY"`
	CreatedAt   string             `json:"created_at" bson:"created_at"`
	UpdatedAt   string             `json:"updated_at" bson:"updated_at"`
	DeleteAt    string             `json:"deleted_at" bson:"deleted_at"`
}

type ProductCategory struct {
}

/************************ Implement setter/getter method for Product ************************/
func (product *Product) GetID() primitive.ObjectID {
	return product.ID
}

func (product *Product) SetID(ID primitive.ObjectID) {
	product.ID = ID
}

func (product *Product) GetProductCode() string {
	return product.ProductCode
}

func (product *Product) SetProductCode(ProductCode string) {
	product.ProductCode = ProductCode
}

func (product *Product) GetName() string {
	return product.Name
}

func (product *Product) SetName(Name string) {
	product.Name = Name
}

func (product *Product) GetCategory() models.Category {
	return product.Category
}

func (product *Product) SetCategory(Category string) {
	product.Category.Name = Category
}

func (product *Product) GetInventory() models.Inventory {
	return product.Inventory
}

func (product *Product) SetInventory(Inventory models.Inventory) {
	product.Inventory = Inventory
}

func (product *Product) GetBrand() models.Brand {
	return product.Brand
}
func (product *Product) SetBrand(Brand models.Brand) {
	product.Brand = Brand
}

func (product *Product) GetSummary() string {
	return product.Summary
}

func (product *Product) SetSummary(Summary string) {
	product.Summary = Summary
}

func (product *Product) GetDescription() string {
	return product.Description
}

func (product *Product) SetDescription(Description string) {
	product.Description = Description
}

func (product *Product) GetImageFile() string {
	return product.ImageFile
}

func (product *Product) SetImageFile(ImageFile string) {
	product.ImageFile = ImageFile
}

func (product *Product) GetDiscount() models.Discount {
	return product.Discount
}

func (product *Product) SetDiscount(Discount models.Discount) {
	product.Discount = Discount
}

func (product *Product) GetPrice() float64 {
	return product.Price
}
func (product *Product) SetPrice(Price float64) {
	product.Price = Price
}

func (product *Product) GetSKU() string {
	return product.SKU
}
func (product *Product) SetSKU(SKU string) {
	product.SKU = SKU
}

func (product *Product) GetCreatedAt() string {
	return product.CreatedAt
}

func (product *Product) SetCreatedAt(CreatedAt string) {
	product.CreatedAt = CreatedAt
}
func (product *Product) GetUpdatedAt() string {
	return product.UpdatedAt
}

func (product *Product) SetUpdatedAt(UpdatedAt string) {
	product.UpdatedAt = UpdatedAt
}

func (product *Product) GetDeleteAt() string {
	return product.DeleteAt
}

func (product *Product) SetDeleteAt(DeleteAt string) {
	product.DeleteAt = DeleteAt
}

/************************ Filter info ************************/
func FilteredResponse(p *Product) Product {
	return Product{
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
