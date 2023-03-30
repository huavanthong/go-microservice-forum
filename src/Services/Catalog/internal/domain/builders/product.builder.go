package builders

import (
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/domain/entities"
)

type ProductBuilder struct {
	id          string
	name        string
	category    string
	summary     string
	description string
	imageFile   string
	price       float64
	productCode string
	sku         string
	createdAt   string
	updatedAt   string
}

func NewProductBuilder() *ProductBuilder {
	return &ProductBuilder{}
}

func (pb *ProductBuilder) SetID(id string) *ProductBuilder {
	pb.id = id
	return pb
}

func (pb *ProductBuilder) SetName(name string) *ProductBuilder {
	pb.name = name
	return pb
}

func (pb *ProductBuilder) SetCategory(category string) *ProductBuilder {
	pb.category = category
	return pb
}

func (pb *ProductBuilder) SetSummary(summary string) *ProductBuilder {
	pb.summary = summary
	return pb
}

func (pb *ProductBuilder) SetDescription(description string) *ProductBuilder {
	pb.description = description
	return pb
}

func (pb *ProductBuilder) SetImageFile(imageFile string) *ProductBuilder {
	pb.imageFile = imageFile
	return pb
}

func (pb *ProductBuilder) SetPrice(price float64) *ProductBuilder {
	pb.price = price
	return pb
}

func (pb *ProductBuilder) SetProductCode(productCode string) *ProductBuilder {
	pb.productCode = productCode
	return pb
}

func (pb *ProductBuilder) SetSKU(sku string) *ProductBuilder {
	pb.sku = sku
	return pb
}

func (pb *ProductBuilder) SetCreatedAt(createdAt string) *ProductBuilder {
	pb.createdAt = createdAt
	return pb
}

func (pb *ProductBuilder) SetUpdatedAt(updatedAt string) *ProductBuilder {
	pb.updatedAt = updatedAt
	return pb
}

func (pb *ProductBuilder) Build() *entities.Product {
	return &entities.Product{
		ID:          pb.id,
		Name:        pb.name,
		Category:    pb.category,
		Summary:     pb.summary,
		Description: pb.description,
		ImageFile:   pb.imageFile,
		Price:       pb.price,
		ProductCode: pb.productCode,
		SKU:         pb.sku,
		CreateAt:    pb.createdAt,
		UpdateAt:    pb.updatedAt,
	}
}
