/*
 * @File: payload.user.req.go
 * @Description: wrapper user request json
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */
package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SignUpInput struct
type SignUpInput struct {
	Name            string `json:"name" bson:"name" binding:"required" example:"John Doe"`
	Email           string `json:"email" bson:"email" binding:"required" example:"johndoe@gmail.com"`
	Password        string `json:"password" bson:"password" binding:"required,min=8" example:"password123"`
	PasswordConfirm string `json:"passwordConfirm" bson:"passwordConfirm,omitempty" binding:"required" example:"password123"`
}

// SignInInput struct
type SignInInput struct {
	Email    string `json:"email" bson:"email" binding:"required" example:"johndoe@gmail.com"`
	Password string `json:"password" bson:"password" binding:"required" example:"password123"`
}

type UpdateDBUser struct {
	ID              primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name            string             `json:"name,omitempty" bson:"name,omitempty"`
	Email           string             `json:"email,omitempty" bson:"email,omitempty"`
	Password        string             `json:"password,omitempty" bson:"password,omitempty"`
	PasswordConfirm string             `json:"passwordConfirm,omitempty" bson:"passwordConfirm,omitempty"`
	Role            string             `json:"role,omitempty" bson:"role,omitempty"`
	Provider        string             `json:"provider" bson:"provider"`
	Photo           string             `json:"photo,omitempty" bson:"photo,omitempty"`
	Verified        bool               `json:"verified,omitempty" bson:"verified,omitempty"`
	CreatedAt       time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt       time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

//  ForgotPasswordInput struct
type ForgotPasswordInput struct {
	Email string `json:"email" binding:"required" example:"johndoe@gmail.com"`
}

//  ResetPasswordInput struct
type ResetPasswordInput struct {
	Password        string `json:"password" binding:"required" example:"password1234"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required" example:"password1234"`
}
