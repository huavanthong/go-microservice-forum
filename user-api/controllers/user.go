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
	"github.com/huavanthong/microservice-golang/user-api/payload"
	"github.com/huavanthong/microservice-golang/user-api/security"
	"github.com/huavanthong/microservice-golang/user-api/utils"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

// Define user manages
type User struct {
	utils   utils.Utils
	userDAO daos.User
}

// Authenticate godoc
// @Summary Check user authentication
// @Description Authenticate user
// @Tags admin
// @Security ApiKeyAuth
// @Accept  multipart/form-data
// @Param user formData string true "Username"
// @Param email formData string true "Email"
// @Param password formData string true "Password"
// @Failure 401 {object} payload.Error
// @Failure 500 {object} payload.Error
// @Success 200 {object} security.Token
// @Router /admin/auth/signin [post]
func (u *User) Authenticate(ctx *gin.Context) {

	// init variable
	var err error

	// get parameter value from request through PostForm
	var accountInfo models.Account
	if err := ctx.ShouldBind(&accountInfo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"binding account info error": err.Error()})
		return
	}

	// var user models.User
	_, err = u.userDAO.Login(accountInfo)

	if err == nil {
		var tokenString string
		// Generate token string
		tokenString, err = u.utils.GenerateJWT(accountInfo.UserName, "")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, payload.Error{common.StatusCodeUnknown, err.Error()})
			log.Debug("[ERROR]: ", err)
			return
		}

		token := security.Token{tokenString}
		// Return token string to the client
		ctx.JSON(http.StatusOK, token)
	} else {
		ctx.JSON(http.StatusUnauthorized, payload.Error{common.StatusCodeUnknown, err.Error()})
	}
}

// AddUser godoc
// @Summary Add a new user
// @Description Add a new user
// @Tags user
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Param user body models.Account true "Add user"
// @Failure 500 {object} payload.Error
// @Failure 400 {object} payload.Error
// @Success 200 {object} payload.Message
// @Router /users [post]
func (u *User) AddUser(ctx *gin.Context) {
	// bind user info to json getting context
	var addAccount models.Account

	// validate data on user data
	// For binding data using go-playground in GIN through models.AddUser
	if err := ctx.ShouldBindJSON(&addAccount); err != nil {
		ctx.JSON(http.StatusInternalServerError, payload.Error{common.StatusCodeUnknown, err.Error()})
		return
	}

	// create user from models
	user := models.User{
		ID:            bson.NewObjectId(),
		Name:          addAccount.UserName,
		Email:         addAccount.Email,
		Password:      addAccount.Password,
		LoginAttempts: []models.LoginAttempt{},
	}

	// insert user to DB
	err := u.userDAO.Insert(user)

	// write response
	if err == nil {
		ctx.JSON(http.StatusOK, payload.Message{"Successfully"})
		log.Debug("Registered a new user = " + user.Name + ", email: " + user.Email + ", password = " + user.Password)
	} else {
		ctx.JSON(http.StatusInternalServerError, payload.Error{common.StatusCodeUnknown, err.Error()})
		log.Debug("[ERROR]: ", err)
	}
}

// ListUsers godoc
// @Summary List all existing users
// @Description List all existing users
// @Tags user
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Failure 500 {object} payload.Error
// @Success 200 {array} models.User
// @Router /users/list [get]
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
		ctx.JSON(http.StatusInternalServerError, payload.Error{common.StatusCodeUnknown, err.Error()})
		log.Debug("[ERROR]: ", err)
	}

}

// GetUserByID godoc
// @Summary Get a user by ID
// @Description Get a user by ID
// @Tags user
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Param id path string true "User ID"
// @Failure 500 {object} payload.Error
// @Success 200 {object} models.User
// @Router /users/detail/{id} [get]
// GetUserByID get a user by id in DB
func (u *User) GetUserByID(ctx *gin.Context) {

	// filter parameter id from context
	id := ctx.Params.ByName("id")

	// find user by id
	user, err := u.userDAO.GetByID(id)

	// write response
	if err == nil {
		ctx.JSON(http.StatusOK, user)
	} else {
		ctx.JSON(http.StatusInternalServerError, payload.Error{common.StatusCodeUnknown, err.Error()})
		fmt.Errorf("[ERROR]: ", err)
	}
}

// GetUserByParams godoc
// @Summary Get a user by ID parameter
// @Description Get a user by ID parameter
// @Tags user
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Param id query string true "User ID"
// @Failure 500 {object} payload.Error
// @Success 200 {object} models.User
// @Router /users [get]
// GetUserByID get a user by id in DB
func (u *User) GetUserByParams(ctx *gin.Context) {

	// filter parameter id from request on context
	id := ctx.Request.URL.Query()["id"][0]

	// find user by id
	user, err := u.userDAO.GetByID(id)

	// write response
	if err == nil {
		ctx.JSON(http.StatusOK, user)
	} else {
		ctx.JSON(http.StatusInternalServerError, payload.Error{common.StatusCodeUnknown, err.Error()})
		fmt.Errorf("[ERROR]: ", err)
	}
}

// DeleteUserByID godoc
// @Summary Delete a user by ID
// @Description Delete a user by ID
// @Tags user
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Param id path string true "User ID"
// @Failure 500 {object} payload.Error
// @Success 200 {object} payload.Message
// @Router /users/{id} [delete]
func (u *User) DeleteUserByID(ctx *gin.Context) {
	// filter parameter id context
	id := ctx.Params.ByName("id")

	// delete user by id
	err := u.userDAO.DeleteByID(id)

	// write response
	if err == nil {
		ctx.JSON(http.StatusOK, payload.Message{"Successfully"})
	} else {
		ctx.JSON(http.StatusInternalServerError, payload.Error{common.StatusCodeUnknown, err.Error()})
		fmt.Errorf("[ERROR]: ", err)
	}
}

// UpdateUser godoc
// @Summary Update an existing user
// @Description Update an existing user
// @Tags user
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Param user body models.User true "User ID"
// @Failure 500 {object} payload.Error
// @Success 200 {object} payload.Message
// @Router /users [patch]
func (u *User) UpdateUser(ctx *gin.Context) {

	// bind user data to json
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, payload.Error{common.StatusCodeUnknown, err.Error()})
		return
	}

	// update user by user
	err := u.userDAO.Update(user)

	// write response
	if err == nil {
		ctx.JSON(http.StatusOK, payload.Message{"Successfully"})
		fmt.Errorf("Update a new user = " + user.Name + ", email: " + user.Email + ", password = " + user.Password)
	} else {
		ctx.JSON(http.StatusInternalServerError, payload.Error{common.StatusCodeUnknown, err.Error()})
		fmt.Errorf("[ERROR]: ", err)
	}
}

// ChangePassword godoc
// @Summary Change password
// @Description Change password of user from the old password
// @Tags user
// @Accept  json
// @Produce  json
// @Param password formData string true "Password"
// @Failure 500 {object} payload.Error
// @Success 200 {object} payload.Message
// @Router /users/changepassword [patch]
func (u *User) ChangePassword(ctx *gin.Context) {

	// get parameter value from request through PostForm
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")

	// change password user
	err := u.userDAO.ChangePassword(email, password)

	// write response
	if err == nil {
		ctx.JSON(http.StatusOK, payload.Message{"Successfully"})
	} else {
		ctx.JSON(http.StatusInternalServerError, payload.Error{common.StatusCodeUnknown, err.Error()})
	}
}
