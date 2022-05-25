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
	ID       bson.ObjectId `bson:"_id" json:"id" example:"5bbdadf782ebac06a695a8e7" `
	Name     string        `bson:"name" json:"name" example:"raycad" `
	Password string        `bson:"password" json:"password" example:"raycad"`
}

// AddUser information
type AddUser struct {
	Name     string `json:"name" example:"User Name"`
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
