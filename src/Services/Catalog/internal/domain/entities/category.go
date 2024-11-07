package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/************************ Define structure product ************************/
type Category struct {
	ID            primitive.ObjectID `json:"id" bson:"_id" example:"5bbdadf782ebac06a695a8e7"`
	CategoryCode  string             `json:"ccode" bson:"ccode" binding:"required" example:"c1000"`
	Name          string             `json:"name" bson:"name" binding:"required" example:"ao-khoac-nu"`
	SubCategories []string           `json:"subcategory" bson:"subcategory"`
	Description   string             `json:"description" bson:"description" example:"Ao khoac thoi trang cho nu"`
	CreatedAt     string             `json:"created_at" bson:"created_at"`
	UpdatedAt     string             `json:"updated_at" bson:"updated_at"`
	DeleteAt      string             `json:"deleted_at" bson:"deleted_at"`
}

type CategoryFilter struct {
	Category string
	Brand    string
}
