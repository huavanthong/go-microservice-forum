/*
 * @File: controllers.user.go
 * @Description: Implements User API logic functions
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */
package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/huavanthong/microservice-golang/user-api/daos"
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

func (u *User) ListUsers(ctx *gin.Context) {

}

func (u *User) GetUserByID(ctx *gin.Context) {

}

func (u *User) GetUserByParams(ctx *gin.Context) {

}

func (u *User) DeleteUserByID(ctx *gin.Context) {

}

func (u *User) UpdateUser(ctx *gin.Context) {

}
