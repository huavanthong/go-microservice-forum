/*
 * @File: main.go
 * @Description: Creates HTTP server & API groups of the UserManagement Service
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */
package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/contrib/jwt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/huavanthong/microservice-golang/user-api/common"
	"github.com/huavanthong/microservice-golang/user-api/controllers"
	"github.com/huavanthong/microservice-golang/user-api/databases"
	"github.com/huavanthong/microservice-golang/user-api/securiy/google"
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
	u := controllers.User{}
	p := controllers.Profile{}

	// generate google token
	token, err := google.RandToken(64)
	if err != nil {
		log.Fatal("unable to generate random token: ", err)
	}

	store := sessions.NewCookieStore([]byte(token))
	store.Options(sessions.Options{
		Path:   "/",
		MaxAge: 86400 * 7,
	})

	m.router.Use(gin.Logger())
	m.router.Use(gin.Recovery())
	m.router.Use(sessions.Sessions("goquestsession", store))
	m.router.Static("/css", "./static/css")
	m.router.Static("/img", "./static/img")
	m.router.LoadHTMLGlob("templates/*")

	// simple group: v1
	v1 := m.router.Group("/api/v1")
	{
		admin := v1.Group("/admin")
		{
			admin.POST("/auth/signin", u.Authenticate)
			// admin.POST("/auth/signin/social", c.AuthSocial)
			// admin.POST("/auth/signin/webadmin", c.AuthWebAdmin)

		}

		user := v1.Group("/users")

		// Todo: Implement APIs need to be authenticated
		user.Use(jwt.Auth(common.Config.JwtSecretPassword))
		{
			user.POST("", u.AddUser)
			user.GET("/list", u.ListUsers)
			user.GET("detail/:id", u.GetUserByID)
			user.GET("/", u.GetUserByParams)
			user.DELETE(":id", u.DeleteUserByID)
			user.PATCH("", u.UpdateUser)
		}

		profile := v1.Group("/profile")
		profile.Use(jwt.Auth(common.Config.JwtSecretPassword))
		{
			profile.POST(":userid", p.AddProfile)
			profile.GET(":userid", p.GetProfileByUserId)
			profile.PUT(":userid", p.UpdateProfileByUserId)
			profile.DELETE(":userid", p.DeteleProfileByUserId)

		}

	}

	m.router.GET("/testserver", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	m.router.Run(common.Config.Port)

}
