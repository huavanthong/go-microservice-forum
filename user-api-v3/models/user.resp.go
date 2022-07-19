/*
 * @File: payload.user.resp.go
 * @Description: wrapper user response json
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */
package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserResponse struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" example:"5bbdadf782ebac06a695a8e7"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty" example:"John Doe"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty" example:"johndoe@gmail.com"`
	Role      string             `json:"role,omitempty" bson:"role,omitempty" example:"user"`
	Photo     string             `json:"photo,omitempty" bson:"photo,omitempty" example:"http://www.golangprograms.com/skin/frontend/base/default/logo.png"`
	Provider  string             `json:"provider" bson:"provider" example:"google oauth2"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at" example: "2022-07-10T12:34:10.91Z"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at" example: "2022-07-10T12:34:10.91Z"`
}

type DBResponse struct {
	ID              primitive.ObjectID `json:"id" bson:"_id"`
	Name            string             `json:"name" bson:"name"`
	Email           string             `json:"email" bson:"email"`
	Password        string             `json:"password" bson:"password"`
	PasswordConfirm string             `json:"passwordConfirm,omitempty" bson:"passwordConfirm,omitempty"`
	Provider        string             `json:"provider" bson:"provider"`
	Photo           string             `json:"photo,omitempty" bson:"photo,omitempty"`
	Role            string             `json:"role" bson:"role"`
	Verified        bool               `json:"verified" bson:"verified"`
	CreatedAt       time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at" bson:"updated_at"`
}

func FilteredResponse(user *DBResponse) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Role:      user.Role,
		Provider:  user.Provider,
		Photo:     user.Photo,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
