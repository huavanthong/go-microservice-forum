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
	UserID        bson.ObjectId  `bson:"_id" json:"userid" example:"5bbdadf782ebac06a695a8e7" `
	UserName      string         `bson:"name" json:"username" example:"hvthong" `
	Email         string         `bson:"email" json:"email" validate:"required,email"`
	Password      string         `bson:"password" json:"password" example:"raycad"`
	CreatedAt     string         `bson:"-" json:"-"`
	UpdateAt      string         `bson:"-" json:"-"`
	LastLoginAt   string         `bson:"-" json:"-"`
	LoginAttempts []LoginAttempt `bson:"-" json:"-"`
	Role          Role           `bson:"-" json:"-"`
	Activated     bool           `bson:"-" json:"-"`
}

// Login is a retrieved and authentiacted user.
type LoginAttempt struct {
	AccountName string `json:"accountname"`
	Password    string `json:"password"`
	IPNumber    string `json:"ipnumber"`
	BrowerType  string `json:"browertype"`
	Success     string `json:"success"`
	CreateDate  string `json:"createdate"`
}

// User Role
type Role struct {
	RoleName string
	RoleNote string
	Actions  []Action
}

// Action for role
type Action struct {
	ActionName string
	ActionURL  string
}

// AddUser information
type AddUser struct {
	Name     string `json:"name" example:"User Name"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" example:"User Password"`
}

// Validate user
func (a *AddUser) Validate() error {

	// check sql injection hacking
	a.CheckSQLInjection(*a)

	switch {
	case len(a.Name) == 0:
		return errors.New(common.ErrNameEmpty)
	case len(a.Password) == 0:
		return errors.New(common.ErrPasswordEmpty)

	default:
		return nil
	}

}

// Check SQL injection hacking
func (a *AddUser) CheckSQLInjection(data AddUser) {

	(*a).Name = security.Santize(data.Name)
	(*a).Password = security.Santize(data.Password)
}
