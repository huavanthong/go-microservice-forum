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
	Category    models.Category    `json:"category" bson:"category" binding:"required,gt=0,lt=255" example:"Phone"`
	Inventory   models.Inventory   `json:"inventory" bson:"inventory" binding:"required" example:"10"`
	Brand       models.Brand       `json:"brand" bson:"brand" binding:"required" example:"Apple"`
	Summary     string             `json:"summary" bson:"summary" binding:"required,gt=0,lt=10000" example:"Iphone 14 Pro Gold"`
	Description string             `json:"description" bson:"description" binding:"required,gt=0,lt=10000" example:"Iphone 14 Pro Gold 256GB"`
	ImageFile   string             `json:"imageFile" bson:"imageFile" binding:"required" example:"default.png"`
	Discount    models.Discount    `json:"discount" bson:"discount" binding:"required" example:"10"`
	Price       float64            `json:"price" bson:"price" binding:"required,min=0.01" example:"1400"`
	SKU         string             `json:"sku" bson:"sku" example:"ABC-XYZ-OXY"`
	CreatedAt   string             `json:"created_at" bson:"created_at"`
	UpdatedAt   string             `json:"updated_at" bson:"updated_at"`
	DeleteAt    string             `json:"deleted_at" bson:"deleted_at"`
}

/************************ Implement setter/getter method for Product ************************/
func (product *Product) GetID() primitive.ObjectID {
	return product.ID
}

func (product *Product) SetID(ID primitive.ObjectID) *Product {
	product.ID = ID
	return product
}

func (product *Product) GetProductCode() string {
	return product.ProductCode
}

func (product *Product) SetProductCode(ProductCode string) *Product {
	product.ProductCode = ProductCode
	return product
}

func (product *Product) GetName() string {
	return product.Name
}

func (product *Product) SetName(Name string) *Product {
	product.Name = Name
	return product
}

func (product *Product) GetCategory() models.Category {
	return product.Category
}

func (product *Product) SetCategories(Category models.Category) *Product {
	product.Category = Category
	return product
}

func (product *Product) GetInventory() models.Inventory {
	return product.Inventory
}

func (product *Product) SetInventory(Inventory models.Inventory) *Product {
	product.Inventory = Inventory
	return product
}

func (product *Product) GetBrand() models.Brand {
	return product.Brand
}
func (product *Product) SetBrand(Brand models.Brand) *Product {
	product.Brand = Brand
	return product
}

func (product *Product) GetSummary() string {
	return product.Summary
}

func (product *Product) SetSummary(Summary string) *Product {
	product.Summary = Summary
	return product
}

func (product *Product) GetDescription() string {
	return product.Description
}

func (product *Product) SetDescription(Description string) *Product {
	product.Description = Description
	return product
}

func (product *Product) GetImageFile() string {
	return product.ImageFile
}

func (product *Product) SetImageFile(ImageFile string) *Product {
	product.ImageFile = ImageFile
	return product
}

func (product *Product) GetDiscount() models.Discount {
	return product.Discount
}
func (product *Product) SetDiscount(Discount models.Discount) *Product {
	product.Discount = Discount
	return product
}

func (product *Product) GetPrice() float64 {
	return product.Price
}
func (product *Product) SetPrice(Price float64) *Product {
	product.Price = Price
	return product
}

func (product *Product) GetSKU() string {
	return product.SKU
}
func (product *Product) SetSKU(SKU string) *Product {
	product.SKU = SKU
	return product
}

func (product *Product) GetCreatedAt() string {
	return product.CreatedAt
}

func (product *Product) SetCreatedAt(CreatedAt string) *Product {
	product.CreatedAt = CreatedAt
	return product
}
func (product *Product) GetUpdatedAt() string {
	return product.UpdatedAt
}

func (product *Product) SetUpdatedAt(UpdatedAt string) *Product {
	product.UpdatedAt = UpdatedAt
	return product
}

func (product *Product) GetDeleteAt() string {
	return product.DeleteAt
}

func (product *Product) SetDeleteAt(DeleteAt string) *Product {
	product.DeleteAt = DeleteAt
	return product
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
