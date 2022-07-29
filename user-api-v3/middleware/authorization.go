package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/casbin/casbin"
	"github.com/zupzup/casbin-http-role-example/model"
)

// Authorizer is a middleware for authorization
func Authorizer(e *casbin.Enforcer, users model.Users) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		session := sessions.Default(ctx)

		role, err := session.Get("role")
		// check role
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Errors()})
			return
		}

		if role == "" {
			role = "anonymous"
		}

		// if it's a member, check if the user still exists
		if role == "member" {
			uid, err := session.Get("userID")
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
				return
			}

			exists := users.Exists(uid)
			if !exists {
				ctx.JSON(http.StatusForbidden, gin.H{"status": "fail", "message": errors.New("user does not exist")})
				return
			}
		}

		// casbin enforce
		method := ctx.Request.Method
		path := ctx.URL.Path
		allowed, err := e.enforcer.Enforce(role, path, method)
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
