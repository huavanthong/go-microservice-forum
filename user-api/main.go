/*
 * @File: main.go
 * @Description: Creates HTTP server & API groups of the UserManagement Service
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */
package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/huavanthong/microservice-golang/user-api/common"
	"github.com/huavanthong/microservice-golang/user-api/controllers"
	"github.com/huavanthong/microservice-golang/user-api/databases"
)

// Main manages main golang application
type Main struct {
	router *gin.Engine
}

func (m *Main) initServer() error {
	var err error
	// Load config file
	err = common.LoadConfig()
	if err != nil {
		return err
	}

	// Initialize User database
	err = databases.Database.Init()
	if err != nil {
		return err
	}

	// Setting Gin logger
	if common.Config.EnableGinFileLog {
		f, _ := os.Create("logs/gin.log")
		if common.Config.EnableGinConsoleLog {
			gin.DefaultWriter = io.MultiWriter(os.Stdout, f)
		} else {
			gin.DefaultWriter = io.MultiWriter(f)
		}
	} else {
		if !common.Config.EnableGinConsoleLog {
			gin.DefaultWriter = io.MultiWriter()
		}
	}

	m.router = gin.Default()

	return nil
}

func main() {

	// init application
	m := Main{}

	// initialize server
	if m.initServer() != nil {
		return
	}
	// init server failed, Close connection to DB
	defer databases.Database.Close()

	// init a controllers
	c := controllers.User{}

	// simple group: v1
	v1 := m.router.Group("/api/v1")
	{
		admin := v1.Group("/admin")
		{
			admin.POST("/auth", c.Authenticate)
			admin.POST("/auth/login", c.Login)
			admin.POST("/auth/register", c.Register)
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

	m.router.GET("/testserver", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	m.router.Run(common.Config.Port)

}
