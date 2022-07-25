package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/huavanthong/microservice-golang/user-api-v3/controllers"
	"github.com/huavanthong/microservice-golang/user-api-v3/services"
)

type AdminRouteController struct {
	adminController controllers.AdminController
}

func NewRouteAdminController(adminController controllers.AdminController) AdminRouteController {
	return AdminRouteController{adminController}
}

func (ac *AdminRouteController) AdminRoute(rg *gin.RouterGroup, adminService services.AdminService) {

	router := rg.Group("admin")
	router.Use(middleware.DeserializeUser(adminService))
	router.GET("/list", ac.adminService.GetAllUsers)
	router.GET("/detail/:userId", ac.adminService.GetUserByID)
	router.GET("/detail/", ac.adminService.GetUserByEmail)
}
