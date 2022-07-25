package controllers

import (
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
// @Param Authorization header string true "Token"
// @Failure 500 {object} payload.Error
// @Success 200 {array} models.User
// @Router /admin/list [get]
// ListUsers get all users exist in DB
func (ac *AdminController) GetAllUsers(ctx *gin.Context) {
	// array of users
	var users []models.User

	// call admin service to get all the exist users
	users, err = ac.adminService.GetAllUsers()
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
		payload.CreatePostSuccess{
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
// @Param Authorization header string true "Token"
// @Param userId path string true "User ID"
// @Failure 500 {object} payload.Response
// @Success 200 {object} models.AdminGetUserSuccess
// @Router /admin/detail/{userId} [get]
// GetUserByID get a user by id in DB
func (ac *AdminController) GetUserByID(ctx *gin.Context) {

	// filter parameter id from context
	id := ctx.Params.ByName("userId")

	// find user by id
	user, err := ac.adminService.GetByID(id)

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
// @Description User find another user by email
// @Tags users
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param email query string true "Email"
// @Failure 404 {object} payload.Response
// @Failure 502 {object} payload.Response
// @Success 200 {object} payload.AdminGetUserSuccess
// @Router /users/ [get]
// SignUp User
func (uc *UserController) GetUserByEmail(ctx *gin.Context) {

	// get user ID from URL path
	email := ctx.Request.URL.Query()["email"][0]

	// call admin service to find user by email
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
			Data:    user,
		})

}
