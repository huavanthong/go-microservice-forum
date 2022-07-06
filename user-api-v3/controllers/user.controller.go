package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wpcodevo/golang-mongodb/models"
	"github.com/wpcodevo/golang-mongodb/services"
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
// @Success 200 {string} http.StatusOK
// @Router /users/me [get]
// SignUp User
func (uc *UserController) GetMe(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(*models.DBResponse)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": models.FilteredResponse(currentUser)}})
}
