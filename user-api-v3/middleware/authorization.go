package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/huavanthong/microservice-golang/user-api-v3/services"

	casbin "github.com/casbin/casbin/v2"
)

// Authorizer is a middleware for authorization
func Authorizer(e *casbin.Enforcer, userService services.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		session := sessions.Default(ctx)
		role := session.Get("role")
		// // check role
		// if err != nil {
		// 	ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Errors()})
		// 	return
		// }
		fmt.Printf("Check 1: ", role)
		if role == "" {
			role = "anonymous"
		}

		// if it's a member, check if the user still exists
		if role == "member" {
			uid := session.Get("userID")
			// if err != nil {
			// 	ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
			// 	return
			// }

			// find user by ids
			user, err := userService.FindUserById(fmt.Sprint(uid))
			fmt.Printf("Check 2: ", user)

			if err != nil {
				ctx.JSON(http.StatusForbidden, gin.H{"status": "fail", "message": errors.New("user does not exist")})
				return
			}
		}

		// casbin enforce
		method := ctx.Request.Method
		path := ctx.Request.URL.Path
		allowed, err := e.Enforce(role, path, method)
		fmt.Printf("Check 3: ", allowed)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		if allowed {
			// pass to next handler
			ctx.Next()
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": errors.New("unauthorized")})
			return
		}
	}
}
