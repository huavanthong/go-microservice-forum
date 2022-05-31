/*
 * @File: main_test.go
 * @Description: Creates test case for HTTP server & API groups of the UserManagement Service
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 * Rereference: https://circleci.com/blog/gin-gonic-testing/
 */
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/huavanthong/microservice-golang/user-api/controllers"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetAllUsers(t *testing.T) {
	// mockResponse := `{"message":"Welcome to the Tech Company listing API with Golang"}`

	// initialize gin router
	r := SetUpRouter()

	c := controllers.User{}

	r.GET("/api/v1/users/list", c.ListUsers)

	req, _ := http.NewRequest("GET", "/api/v1/users/list", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	fmt.Println(responseData)
	assert.Equal(t, http.StatusOK, w.Code)
}
