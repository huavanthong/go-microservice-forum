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
