package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// understanding json encoding, [here](https://pkg.go.dev/encoding/json#Marshal)
type User struct {
	ID              primitive.ObjectID `bson:"_id" json:"id" example:"5bbdadf782ebac06a695a8e7" `
	Name            string             `json:"name" bson:"name" binding:"required" example:"John Doe"`
	Email           string             `json:"email" bson:"email" binding:"required" example:"johndoe@gmail.com"`
	Password        string             `json:"password" bson:"password" binding:"required,min=8" example:"password123"`
	PasswordConfirm string             `json:"passwordConfirm" bson:"passwordConfirm,omitempty" binding:"required" example:"password123"`
	Role            string             `json:"role" bson:"role"`
	Provider        string             `json:"provider,omitempty" bson:"provider,omitempty"`
	Photo           string             `json:"photo,omitempty" bson:"photo,omitempty"`
	Verified        bool               `json:"verified" bson:"verified"`
	LoginAttempts   []LoginAttempt     `bson:"loginattempts" json:"loginattempts"`
	Activated       bool               `bson:"activated" json:"activated"`
	CreatedAt       time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at" bson:"updated_at"`
	LastLoginAt     string             `json:"lastloginat" bson:"lastloginat"`
}

// Login is a retrieved and authentiacted user.
type LoginAttempt struct {
	AccountName string `bson:"accountname" json:"accountname"` // define account name is not correct with user id
	Password    string `bson:"password" json:"password"`
	IPNumber    string `bson:"ipnumber" json:"ipnumber"`
	BrowerType  string `bson:"browertype" json:"browertype"`
	Success     string `bson:"success" json:"success"`
	CreateDate  string `bson:"createdate" json:"createdate"`
}
