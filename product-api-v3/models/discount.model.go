package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/************************ Define structure product ************************/
type Discount struct {
	ID              primitive.ObjectID `bson:"_id" json:"id" example:"5bbdadf782ebac06a695a8e7"`
	Name            string
	Description     string
	DiscountPercent float32
	Active          bool
	CreatedAt       string `json:"created_at" bson:"created_at"`
	UpdatedAt       string `json:"updated_at" bson:"updated_at"`
	DeleteAt        string `json:"deleted_at" bson:"deleted_at"`
}
