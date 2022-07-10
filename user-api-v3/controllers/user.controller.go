package controllers

import (
	"net/http"

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
