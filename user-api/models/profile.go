/*
 * @File: models.profiles.go
 * @Description: Define profile for user
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */

package models

import "gopkg.in/mgo.v2/bson"

// User contains user information
type Profile struct {
	ID             bson.ObjectId `bson:"_id" json:"id" example:"5bbdadf782ebac06a695a8e7"`
	ProfileName    string        `bson:"profilename" json:"profilename"`
	FirstName      string        `bson:"firstname" json:"firstname" binding:"required" example:"John"`
	LastName       string        `bson:"lastname" json:"lastname" example:"Switch"`
	Email          string        `bson:"email" json:"email" binding:"required,email" example:"example@gmail.com"`
	EmailVerified  bool          `bson:"emailverified" json:"emailverified" example:"true"`
	UserID         bson.ObjectId `bson:"_userid" json:"userid" binding:"required" example:"5bbdadf782ebac06a695a8e7"`
	Age            uint8         `bson:"age" json:"age" binding:"omitempty,gte=0,lte=130" example:"30"`
	Gender         string        `bson:"gender" json:"gender" example:"male"`
	PhoneNumber    string        `bson:"phonenumber" json:"phonenumber"`
	Picture        string        `bson:"picture" json:"picture" example:"link to picture"`
	DefaultProfile string        `bson:"defaultprofile" json:"defaultprofile"`
	FavouriteColor string        `bson:"favouritecolor" json:"favouritecolor" binding:"iscolor" example:"#0003"` // alias for 'hexcolor|rgb|rgba|hsl|hsla'
	Addresses      []*Address    `bson:"addresses" json"addresses"`                                              // a person can have a home and cottage...
}

type Contact struct {
	HomePhone     string
	ColleagePhone string
}

// Address houses a users address information
type Address struct {
	Street   string `bson:"street" json:"street"`
	Ward     string `bson:"ward" json:"ward"`
	District string `bson:"district" json:"district"`
	City     string `bson:"city" json:"city"`
	Country  string `bson:"country" json:"country"`
	Zip      string `bson:"zip" json:"zip"`
}
