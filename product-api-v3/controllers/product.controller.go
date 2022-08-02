package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/huavanthong/microservice-golang/product-api-v3/models"
	"github.com/huavanthong/microservice-golang/product-api-v3/payload"
	"github.com/huavanthong/microservice-golang/product-api-v3/services"
)

type ProductController struct {
	logger         zap.Logger
	productService services.ProductService
}

func NewProductController(logger zap.Logger, productService services.ProductService) ProductController {
	return ProductController{logger, productService}
}

// GetAllProducts godoc
// @Summary List all existing products
// @Description List all existing products of store
// @Tags products
// @Accept  json
// @Produce  json
// @Param page path string true "Page"
// @Param limit path string true "Limit"
// @Param currency path string true "Limit"
// @Failure 500 {object} payload.Response
// @Success 200 {array} models.User
// @Router /products [get]
func (pc *ProductController) GetAllProducts(ctx *gin.Context) {

	// get parameter from client
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")
	var currency = ctx.DefaultQuery("currency", "USD")

	intPage, err := strconv.Atoi(page)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	// call product service to get all the exist product in store
	products, err := pc.productService.FindAllProducts(intPage, intLimit, currency)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusOK,
		payload.GetAllProductSuccess{
			Status:  "success",
			Code:    http.StatusCreated,
			Message: "Get all products success",
			Data:    products,
		})
}

// GetProductByID godoc
// @Summary Get a product by ID
// @Description Get a product by ID
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path string true "Product ID"
// @Failure 404 {object} payload.Response
// @Failure 502 {object} payload.Response
// @Success 200 {object} payload.AdminGetUserSuccess
// @Router /products/{id} [get]
// GetUserByID get a user by id in DB
func (pc *ProductController) GetProductByID(ctx *gin.Context) {

	// filter parameter id from context
	productId := ctx.Params.ByName("id")

	// find user by id
	user, err := pc.productService.FindProductByID(productId)

	// catch error
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			ctx.JSON(http.StatusNotFound,
				payload.Response{
					Status:  "fail",
					Code:    http.StatusNotFound,
					Message: err.Error(),
				})
			return
		}
		ctx.JSON(http.StatusBadGateway,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusBadGateway,
				Message: err.Error(),
			})
		return
	}

	// succes
	ctx.JSON(http.StatusOK,
		payload.AdminGetUserSuccess{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Admin get user success",
			Data:    models.AdminFilteredResponse(user),
		})

}

// GetProductByName godoc
// @Summary Get user by Email
// @Description Admin get user info by email
// @Tags admin
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param name query string true "Name of Product"
// @Failure 404 {object} payload.Response
// @Failure 502 {object} payload.Response
// @Success 200 {object} payload.AdminGetUserSuccess
// @Router /admin/ [get]
func (pc *ProductController) GetProductByName(ctx *gin.Context) {

	// get user ID from URL path
	email := ctx.Request.URL.Query()["name"][0]

	// call admin service to get user by email
	user, err := pc.productService.GetUserByEmail(email)
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			ctx.JSON(http.StatusNotFound,
				payload.Response{
					Status:  "fail",
					Code:    http.StatusNotFound,
					Message: err.Error(),
				})
			return
		}
		ctx.JSON(http.StatusBadGateway,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusBadGateway,
				Message: err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusOK,
		payload.AdminGetUserSuccess{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Get user success",
			Data:    models.AdminFilteredResponse(user),
		})

}

// GetUserByEmail godoc
// @Summary Get user by Email
// @Description Admin get user info by email
// @Tags admin
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param email query string true "Email"
// @Failure 404 {object} payload.Response
// @Failure 502 {object} payload.Response
// @Success 200 {object} payload.AdminGetUserSuccess
// @Router /admin/ [get]
func (pc *ProductController) GetProductByCategory(ctx *gin.Context) {

	// get user ID from URL path
	email := ctx.Request.URL.Query()["email"][0]

	// call admin service to get user by email
	user, err := ac.adminService.GetUserByEmail(email)
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			ctx.JSON(http.StatusNotFound,
				payload.Response{
					Status:  "fail",
					Code:    http.StatusNotFound,
					Message: err.Error(),
				})
			return
		}
		ctx.JSON(http.StatusBadGateway,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusBadGateway,
				Message: err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusOK,
		payload.AdminGetUserSuccess{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Get user success",
			Data:    models.AdminFilteredResponse(user),
		})

}

// GetUserByEmail godoc
// @Summary Get user by Email
// @Description Admin get user info by email
// @Tags admin
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param email query string true "Email"
// @Failure 404 {object} payload.Response
// @Failure 502 {object} payload.Response
// @Success 200 {object} payload.AdminGetUserSuccess
// @Router /admin/ [get]
func (pc *ProductController) AddProduct(ctx *gin.Context) {

	// get user ID from URL path
	email := ctx.Request.URL.Query()["email"][0]

	// call admin service to get user by email
	user, err := ac.adminService.GetUserByEmail(email)
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			ctx.JSON(http.StatusNotFound,
				payload.Response{
					Status:  "fail",
					Code:    http.StatusNotFound,
					Message: err.Error(),
				})
			return
		}
		ctx.JSON(http.StatusBadGateway,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusBadGateway,
				Message: err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusOK,
		payload.AdminGetUserSuccess{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Get user success",
			Data:    models.AdminFilteredResponse(user),
		})

}

// UpdatePost godoc
// @Summary Update a post
// @Description User update the exist post
// @Tags posts
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param post body models.UpdatePost true "Update post"
// @Param postId path string true "Post ID"
// @Failure 404 {object} payload.Response
// @Failure 502 {object} payload.Response
// @Success 200 {object} payload.Response
// @Router /posts/{postId} [patch]
func (pc *ProductController) UpdateProduct(ctx *gin.Context) {

	// get post ID from URL path
	postId := ctx.Param("postId")

	// from context, bind a new post info to json
	var post *models.UpdatePost
	if err := ctx.ShouldBindJSON(&post); err != nil {
		ctx.JSON(http.StatusBadGateway,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusBadGateway,
				Message: err.Error(),
			})
		return
	}

	// call post service to update info
	updatedPost, err := pc.postService.UpdatePost(postId, post)
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			ctx.JSON(http.StatusNotFound,
				payload.Response{
					Status:  "fail",
					Code:    http.StatusNotFound,
					Message: err.Error(),
				})
			return
		}
		ctx.JSON(http.StatusBadGateway,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusBadGateway,
				Message: err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusOK,
		payload.UpdatePostSuccess{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Update a exist post success",
			Data:    models.FilteredPostResponse(updatedPost),
		})
}

// DeleteProduct godoc
// @Summary Delete a post
// @Description User delete the post by postId
// @Tags posts
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param postId path string true "Post ID"
// @Failure 404 {object} payload.Response
// @Failure 502 {object} payload.Response
// @Success 204 {object} payload.Response
// @Router /posts/{postId} [delete]
func (pc *ProductController) DeleteProductByID(ctx *gin.Context) {
	// get post ID from URL path
	postId := ctx.Param("postId")

	// call post service to delete post by ID
	err := pc.postService.DeletePost(postId)

	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			ctx.JSON(http.StatusNotFound,
				payload.Response{
					Status:  "fail",
					Code:    http.StatusNotFound,
					Message: err.Error(),
				})
			return
		}
		ctx.JSON(http.StatusBadGateway,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusBadGateway,
				Message: err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusNoContent,
		payload.Response{
			Status:  "success",
			Code:    http.StatusNoContent,
			Message: "Delete post success",
		})
}
