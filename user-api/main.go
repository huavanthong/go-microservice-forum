/*
 * @File: main.go
 * @Description: Creates HTTP server & API groups of the UserManagement Service
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/users", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("/users", func(c *gin.Context) {
		c.String(http.StatusOK, "post a user")
	})

	r.PATCH("/users", func(c *gin.Context) {
		c.String(http.StatusOK, "update a user")
	})

	r.GET("/users/detail/:id", func(c *gin.Context) {
		c.String(http.StatusOK, "get a specific user")
	})

	r.GET("/users/list", func(c *gin.Context) {
		c.String(http.StatusOK, "get a list of user")
	})

	r.DELETE("/users/:id", func(c *gin.Context) {
		c.String(http.StatusOK, "delete a user")
	})

	r.Run(":8080")
}
