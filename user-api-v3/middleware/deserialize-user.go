package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/huavanthong/microservice-golang/user-api-v3/config"
	"github.com/huavanthong/microservice-golang/user-api-v3/services"
	"github.com/huavanthong/microservice-golang/user-api-v3/utils"
)

func DeserializeUser(userService services.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get access token from cache cookie
		var access_token string
		cookie, err := ctx.Cookie("access_token")

		// check authorize
		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			access_token = fields[1]
		} else if err == nil {
			access_token = cookie
		}

		// check user login or not
		if access_token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}

		// get access token publickey to validate jwt token.
		config, _ := config.LoadConfig(".")
		sub, err := utils.ValidateToken(access_token, config.AccessTokenPublicKey)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		// find user by id
		user, err := userService.FindUserById(fmt.Sprint(sub))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "The user belonging to this token no logger exists"})
			return
		}

		// set current user to cookie
		ctx.Set("currentUser", user)

		// pass to next handler
		ctx.Next()
	}
}

// func DeserializeAdmin(adminService services.AdminService) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		var access_token string
// 		cookie, err := ctx.Cookie("access_token")

// 		authorizationHeader := ctx.Request.Header.Get("Authorization")
// 		fields := strings.Fields(authorizationHeader)

// 		if len(fields) != 0 && fields[0] == "Bearer" {
// 			access_token = fields[1]
// 		} else if err == nil {
// 			access_token = cookie
// 		}

// 		if access_token == "" {
// 			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
// 			return
// 		}

// 		config, _ := config.LoadConfig(".")
// 		sub, err := utils.ValidateToken(access_token, config.AccessTokenPublicKey)
// 		if err != nil {
// 			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
// 			return
// 		}

// 		user, err := adminService.FindUserById(fmt.Sprint(sub))
// 		if err != nil {
// 			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "The user belonging to this token no logger exists"})
// 			return
// 		}

// 		ctx.Set("currentUser", user)
// 		ctx.Next()
// 	}
// }
