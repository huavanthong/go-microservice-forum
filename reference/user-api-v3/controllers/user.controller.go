package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/huavanthong/microservice-golang/user-api-v3/models"
	"github.com/huavanthong/microservice-golang/user-api-v3/payload"
	"github.com/huavanthong/microservice-golang/user-api-v3/services"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return UserController{userService}
}

// GetMe godoc
// @Summary Get the current user info
// @Description Get the current user info
// @Tags users
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Success 200 {object} payload.UserRegisterSuccess
// @Router /users/me [get]
// SignUp User
func (uc *UserController) GetMe(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(*models.DBResponse)

	ctx.JSON(http.StatusOK,
		payload.UserRegisterSuccess{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Get the current user info",
			Data:    models.FilteredResponse(currentUser),
		})

}

// GetUserByID godoc
// @Summary Get user by ID
// @Description User find another user by ID
// @Tags users
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param userId path string true "User ID"
// @Failure 404 {object} payload.Response
// @Failure 502 {object} payload.Response
// @Success 200 {object} payload.GetUserSuccess
// @Router /users/{userId}} [get]
// SignUp User
func (uc *UserController) GetUserByID(ctx *gin.Context) {

	// get user ID from URL path
	userId := ctx.Param("userId")

	// call post service to find post by ID
	user, err := uc.userService.FindUserById(userId)
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			ctx.JSON(http.StatusNotFound,
				payload.Response{
					Status:  "fail",
					Code:    http.StatusNotFound,
					Message: err.Error(),
				})
			return
		}
		ctx.JSON(http.StatusBadGateway,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusBadGateway,
				Message: err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusOK,
		payload.GetUserSuccess{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Get user success",
			Data:    models.FilteredResponse(user),
		})

}

// GetUserByEmail godoc
// @Summary Get user by Email
// @Description User find another user by email
// @Tags users
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param email query string true "Email"
// @Failure 404 {object} payload.Response
// @Failure 502 {object} payload.Response
// @Success 200 {object} payload.GetUserSuccess
// @Router /users/ [get]
// SignUp User
func (uc *UserController) GetUserByEmail(ctx *gin.Context) {

	// get user ID from URL path
	email := ctx.Request.URL.Query()["email"][0]

	// call user service to find user by email
	user, err := uc.userService.FindUserByEmail(email)
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			ctx.JSON(http.StatusNotFound,
				payload.Response{
					Status:  "fail",
					Code:    http.StatusNotFound,
					Message: err.Error(),
				})
			return
		}
		ctx.JSON(http.StatusBadGateway,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusBadGateway,
				Message: err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusOK,
		payload.GetUserSuccess{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Get user success",
			Data:    models.FilteredResponse(user),
		})

}
