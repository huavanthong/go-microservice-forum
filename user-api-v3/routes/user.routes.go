package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/huavanthong/microservice-golang/user-api-v3/controllers"
	"github.com/huavanthong/microservice-golang/user-api-v3/middleware"
	"github.com/huavanthong/microservice-golang/user-api-v3/services"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewRouteUserController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(rg *gin.RouterGroup, userService services.UserService) {

	router := rg.Group("users")
	router.Use(middleware.DeserializeUser(userService))
	router.GET("/me", uc.userController.GetMe)
	router.GET("/:userId", uc.userController.GetUserByID)
	router.GET("/:email", uc.userController.GetUserByEmail)
}
