package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/huavanthong/microservice-golang/user-api-v3/payload"
	"github.com/huavanthong/microservice-golang/user-api-v3/services"
)

type AdminController struct {
	adminService services.AdminService
}

func NewAdminController(adminService services.AdminService) AdminController {
	return AdminController{adminService}
}

// GetAllUsers godoc
// @Summary List all existing users
// @Description List all existing users
// @Tags admin
// @Accept  json
// @Produce  json
// @Param page path string true "Post ID"
// @Param limit path string true "Post ID"
// @Failure 500 {object} payload.Response
// @Success 200 {array} models.User
// @Router /admin/list [get]
// ListUsers get all users exist in DB
func (ac *AdminController) GetAllUsers(ctx *gin.Context) {

	// get parameter from client
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, err := strconv.Atoi(page)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	// array of users
	// var users []models.User

	// call admin service to get all the exist users
	users, err := ac.adminService.GetAllUsers(intPage, intLimit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusOK,
		payload.AdminGetAllUserSuccess{
			Status:  "success",
			Code:    http.StatusCreated,
			Message: "Get all users success",
			Data:    users,
		})
}

// GetUserByID godoc
// @Summary Get a user by ID
// @Description Get a user by ID
// @Tags admin
// @Accept  json
// @Produce  json
// @Param userId path string true "User ID"
// @Failure 500 {object} payload.Response
// @Success 200 {object} payload.AdminGetUserSuccess
// @Router /admin/detail/{userId} [get]
// GetUserByID get a user by id in DB
func (ac *AdminController) GetUserByID(ctx *gin.Context) {

	// filter parameter id from context
	userId := ctx.Params.ByName("userId")

	// find user by id
	user, err := ac.adminService.GetUserByID(userId)

	// catch error
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

	// succes
	ctx.JSON(http.StatusOK,
		payload.AdminGetUserSuccess{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Admin get user success",
			Data:    user,
		})

}

// GetUserByEmail godoc
// @Summary Get user by Email
// @Description Admin get user info by email
// @Tags admin
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param email query string true "Email"
// @Failure 404 {object} payload.Response
// @Failure 502 {object} payload.Response
// @Success 200 {object} payload.AdminGetUserSuccess
// @Router /users/ [get]
// SignUp User
func (ac *AdminController) GetUserByEmail(ctx *gin.Context) {

	// get user ID from URL path
	email := ctx.Request.URL.Query()["email"][0]

	// call admin service to get user by email
	user, err := ac.adminService.GetUserByEmail(email)
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
		payload.AdminGetUserSuccess{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Get user success",
			Data:    user,
		})

}
