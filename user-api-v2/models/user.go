/*
 * @File: models.user.go
 * @Description: Defines User model
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */
package models

import (
	"errors"

	"github.com/huavanthong/microservice-golang/user-api/common"
	"github.com/huavanthong/microservice-golang/user-api/security"

	"gopkg.in/mgo.v2/bson"
)

// User information
type User struct {
	ID            bson.ObjectId  `bson:"_id" json:"id" example:"5bbdadf782ebac06a695a8e7" `
	Name          string         `bson:"name" json:"username" example:"hvthong" `
	Email         string         `bson:"email" json:"email" validate:"required,email"`
	Password      string         `bson:"password" json:"password" example:"raycad"`
	CreatedAt     string         `bson:"createat" json:"-"`
	UpdateAt      string         `bson:"updateat" json:"-"`
	LastLoginAt   string         `bson:"lastloginat" json:"-"`
	LoginAttempts []LoginAttempt `bson:"loginattempts" json:"loginattempts"`
	Role          Role           `bson:"role" json:"role"`
	Activated     bool           `bson:"activated" json:"activated"`
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

// User Role
type Role struct {
	RoleName string   `bson:"rolename" json:"rolename"`
	RoleNote string   `bson:"rolenote" json:"rolenote"`
	Actions  []Action `bson:"actions" json:"actions"`
}

// Action for role
type Action struct {
	ActionName string `bson:"actionname" json:"actionname"`
	ActionURL  string `bson:"actionurl" json:"actionurl"`
}

// Account information
type Account struct {
	UserName string `form:"username" binding:"required,max=64" example:"vanthong"`
	Email    string `form:"email" binding:"required,email,max=255" example:"hvthong@gmail.com"`
	Password string `form:"password" binding:"required" example:"User Password"`
}

// Validate user
func (a *Account) Validate() error {

	// check sql injection hacking
	a.CheckSQLInjection(*a)

	switch {
	case len(a.UserName) == 0:
		return errors.New(common.ErrNameEmpty)
	case len(a.Password) == 0:
		return errors.New(common.ErrPasswordEmpty)

	default:
		return nil
	}

}

// Check SQL injection hacking
func (a *Account) CheckSQLInjection(data Account) {

	(*a).UserName = security.Santize(data.UserName)
	(*a).Password = security.Santize(data.Password)
}
