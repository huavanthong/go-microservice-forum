package routes

import (
	casbin "github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/huavanthong/microservice-golang/user-api-v3/controllers"
	"github.com/huavanthong/microservice-golang/user-api-v3/middleware"
	"github.com/huavanthong/microservice-golang/user-api-v3/services"
)

type AuthRouteController struct {
	authController controllers.AuthController
}

func NewAuthRouteController(authController controllers.AuthController) AuthRouteController {
	return AuthRouteController{authController}
}

func (rc *AuthRouteController) AuthRoute(rg *gin.RouterGroup, userService services.UserService, enforcer *casbin.Enforcer) {

	router := rg.Group("/auth")

	router.POST("/register", rc.authController.SignUpUser)
	router.GET("/verifyemail/:verificationCode", rc.authController.VerifyEmail)

	router.Use(middleware.Authorizer(enforcer, userService))
	router.POST("/login", rc.authController.SignInUser)
	router.GET("/refresh", rc.authController.RefreshAccessToken)
	router.GET("/logout", middleware.DeserializeUser(userService), rc.authController.LogoutUser)
	router.POST("/forgotpassword", rc.authController.ForgotPassword)
	router.PATCH("/resetpassword/:resetToken", rc.authController.ResetPassword)
}
