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

// // GetAllUsers godoc
// // @Summary List all existing users
// // @Description List all existing users
// // @Tags admin
// // @Accept  json
// // @Produce  json
// // @Param Authorization header string true "Token"
// // @Failure 500 {object} payload.Error
// // @Success 200 {array} models.User
// // @Router /users/list [get]
// // ListUsers get all users exist in DB
// func (ac *AdminController) GetAllUsers(ctx *gin.Context) {
// 	// array of users
// 	var users []models.User

// 	// call admin service to get all the exist users
// 	users, err = ac.adminService.GetAllUsers()
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError,
// 			payload.Response{
// 				Status:  "fail",
// 				Code:    http.StatusInternalServerError,
// 				Message: err.Error(),
// 			})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK,
// 		payload.CreatePostSuccess{
// 			Status:  "success",
// 			Code:    http.StatusCreated,
// 			Message: "Get all users success",
// 			Data:    users,
// 		})
// }

// // GetUserByID godoc
// // @Summary Get a user by ID
// // @Description Get a user by ID
// // @Tags admin
// // @Accept  json
// // @Produce  json
// // @Param Authorization header string true "Token"
// // @Param id path string true "User ID"
// // @Failure 500 {object} payload.Error
// // @Success 200 {object} models.User
// // @Router /users/detail/{id} [get]
// // GetUserByID get a user by id in DB
// func (u *User) GetUserByID(ctx *gin.Context) {

// 	// filter parameter id from context
// 	id := ctx.Params.ByName("id")

// 	// find user by id
// 	user, err := u.userDAO.GetByID(id)

// 	// write response
// 	if err == nil {
// 		ctx.JSON(http.StatusOK, user)
// 	} else {
// 		ctx.JSON(http.StatusInternalServerError, payload.Error{common.StatusCodeUnknown, err.Error()})
// 		fmt.Errorf("[ERROR]: ", err)
// 	}
// }
