package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreatePostRequest struct {
	Title     string    `json:"title" bson:"title" binding:"required" example:"My post"`
	Content   string    `json:"content" bson:"content" binding:"required" example:"The post tutorial with Golang"`
	Image     string    `json:"image,omitempty" bson:"image,omitempty" example:"default.png"`
	User      string    `json:"user" bson:"user" binding:"required" example:"5bbdadf782ebac06a695a8e7"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type UpdatePost struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title,omitempty" bson:"title,omitempty" example:"My post update"`
	Content   string             `json:"content,omitempty" bson:"content,omitempty" example:"The post tutorial with Golang + Gin"`
	Image     string             `json:"image,omitempty" bson:"image,omitempty" example:"default.png"`
	User      string             `json:"user,omitempty" bson:"user,omitempty" example:"5bbdadf782ebac06a695a8e7"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
