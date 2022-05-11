/*
 * @File: controllers.user.go
 * @Description: Implements User API logic functions
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */
package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huavanthong/microservice-golang/user-api/common"
	"github.com/huavanthong/microservice-golang/user-api/daos"
	"github.com/huavanthong/microservice-golang/user-api/models"
	"github.com/huavanthong/microservice-golang/user-api/utils"
)

// Define user manages
type User struct {
	utils   utils.Utils
	userDAO daos.User
}

func (u *User) Authenticate(ctx *gin.Context) {

}

func (u *User) AddUser(ctx *gin.Context) {

}

// ListUsers get all users exist in DB
func (u *User) ListUsers(ctx *gin.Context) {

	// array of users
	var users []models.User

	// get all users
	users, err := u.userDAO.GetAll()

	// write response
	if err == nil {
		ctx.JSON(http.StatusOK, users)
	} else {
		ctx.JSON(http.StatusInternalServerError, models.Error{common.StatusCodeUnknown, err.Error()})
		fmt.Errorf("[ERROR]: ", err)
	}

}

func (u *User) GetUserByID(ctx *gin.Context) {

	// filter parameter id from context
	id := ctx.Params.ByName("id")

	// find user by id
	user, err := u.userDAO.GetByID(id)

	// write response
	if err == nil {
		ctx.JSON(http.StatusOK, user)
	} else {
		ctx.JSON(http.StatusInternalServerError, models.Error{common.StatusCodeUnknown, err.Error()})
		fmt.Errorf("[ERROR]: ", err)
	}
}

func (u *User) GetUserByParams(ctx *gin.Context) {

	// filter parameter id from request on context
	id := ctx.Request.URL.Query()["id"][0]

	// find user by id
	user, err := u.userDAO.GetByID(id)

	// write response
	if err == nil {
		ctx.JSON(http.StatusOK, user)
	} else {
		ctx.JSON(http.StatusInternalServerError, models.Error{common.StatusCodeUnknown, err.Error()})
		fmt.Errorf("[ERROR]: ", err)
	}
}

func (u *User) DeleteUserByID(ctx *gin.Context) {

}

func (u *User) UpdateUser(ctx *gin.Context) {

}
