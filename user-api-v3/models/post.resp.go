/*
 * @File: payload.post.resp.go
 * @Description: wrapper user response json
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */
package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostResponse struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title,omitempty" bson:"title,omitempty"`
	Content   string             `json:"content,omitempty" bson:"content,omitempty"`
	Image     string             `json:"image,omitempty" bson:"image,omitempty"`
	User      string             `json:"user,omitempty" bson:"user,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

func FilteredResponse(post *DBPost) PostResponse {
	return PostResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		Image:     post.Image,
		User:      post.User,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
}
