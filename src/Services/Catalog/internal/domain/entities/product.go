package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/************************ Define structure product ************************/
type Product struct {
	ID          primitive.ObjectID `bson:"_id" json:"id" example:"5bbdadf782ebac06a695a8e7"`
	ProductCode string             `bson:"pcode" json:"pcode" binding:"required" example:"p123456789"`
	ProductType string             `bson:"ptype" json:"ptype" binding:"required" example:"phone"`
	Name        string             `json:"name" bson:"name" binding:"required,gt=0,lt=255" example:"Iphone 14 Pro"`
	Category    Category           `json:"category" bson:"category" binding:"required,gt=0,lt=255"`
	Inventory   Inventory          `json:"inventory" bson:"inventory" binding:"required"`
	Brand       Brand              `json:"brand" bson:"brand" binding:"required"`
	Summary     string             `json:"summary" bson:"summary" binding:"required,gt=0,lt=10000" example:"Iphone 14 Pro Gold"`
	Description string             `json:"description" bson:"description" binding:"required,gt=0,lt=10000" example:"Iphone 14 Pro Gold 256GB"`
	ImageFile   string             `json:"imageFile" bson:"imageFile" binding:"required" example:"default.png"`
	Discount    Discount           `json:"discount" bson:"discount" binding:"required"`
	Price       float64            `json:"price" bson:"price" binding:"required,min=0.01" example:"1400"`
	SKU         string             `json:"sku" bson:"sku" example:"ABC-XYZ-OXY"`
	CreatedAt   string             `json:"created_at" bson:"created_at"`
	UpdatedAt   string             `json:"updated_at" bson:"updated_at"`
	DeleteAt    string             `json:"deleted_at" bson:"deleted_at"`
}

type ProductFilter struct {
	Category    string
	ProductType string
	Brand       string
}

type Pagination struct {
	Page  int
	Limit int
}

func (p *Pagination) Offset() int {
	return (p.Page - 1) * p.Limit
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
