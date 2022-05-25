/*
 * @File: models.profiles.go
 * @Description: Define profile for user
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */

package models

import "gopkg.in/mgo.v2/bson"

// User contains user information
type Profiles struct {
	ProfileID      bson.ObjectId `bson:"profileid" json:"id" validate:"required,gt=0" example:"5bbdadf782ebac06a695a8e7"`
	ProfileName    string        `bson:"profilename" json:"profilename"`
	FirstName      string        `bson:"firstname" json:"firstname" validate:"required"`
	LastName       string        `bson:"lastname" json:"lastname" validate:"required"`
	Email          string        `bson:"email" json:"email" validate:"required,email"`
	AccountID      int           `bson:"accountid" json:"accountid" validate:"required"`
	Age            uint8         `bson:"age" json:"age" validate:"gte=0,lte=130"`
	PhoneNumber    string        `bson:"phonenumber" json:"phonenumber" validate:"required"`
	DefaultProfile string        `bson:"defaultprofile" json:"defaultprofile"`
	FavouriteColor string        `bson:"favouritecolor" json:"favouritecolor" validate:"iscolor" example:"#000-"` // alias for 'hexcolor|rgb|rgba|hsl|hsla'
	Addresses      []*Address    `bson:"addresses" json"addresses" validate:"required,dive,required"`             // a person can have a home and cottage...
}

// Address houses a users address information
type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
}
