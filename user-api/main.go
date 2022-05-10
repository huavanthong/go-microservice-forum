/*
 * @File: main.go
 * @Description: Creates HTTP server & API groups of the UserManagement Service
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */
package main

import (
	"net/http"

	"./controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	// init a router from gin
	r := gin.Default()

	// init a controllers
	c := controllers.User{}

	// simple group: v1
	v1 := gin.router.Group("/api/v1")
	{
		admin := v1.Group("/admin")
		{
			admin.POST("/auth", c.Authenticate)
		}

		user := v1.Group("/users")

		// Todo: Implement APIs need to be authenticated
		user.POST("", c.AddUser)
		user.GET("/list", c.ListUsers)
		user.GET("detail/:id", c.GetUserByID)
		user.GET("/", c.GetUserByParams)
		user.DELETE(":id", c.DeleteUserByID)
		user.PATCH("", c.UpdateUser)

	}

	r.GET("/testserver", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.Run(":8080")
}
