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
	log            *zap.Logger
	productService services.ProductService
}

func NewProductController(log *zap.Logger, productService services.ProductService) ProductController {
	return ProductController{log, productService}
}

// GetAllProducts godoc
// @Summary List all existing products
// @Description List all existing products of store
// @Tags products
// @Accept  json
// @Produce  json
// @Param page path string true "Page"
// @Param limit path string true "Limit"
// @Param currency path string true "Currency"
// @Failure 500 {object} payload.Response
// @Success 200 {array} models.Product
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
// @Param currency path string true "Currency"
// @Failure 404 {object} payload.Response
// @Failure 502 {object} payload.Response
// @Success 200 {object} payload.GetProductSuccess
// @Router /products/{id} [get]
// GetUserByID get a user by id in DB
func (pc *ProductController) GetProductByID(ctx *gin.Context) {

	// filter parameter id from context
	productId := ctx.Params.ByName("id")

	var currency = ctx.DefaultQuery("currency", "USD")

	// find user by id
	product, err := pc.productService.FindProductByID(productId, currency)

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
		payload.GetProductSuccess{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Get product success",
			Data:    product,
		})

}

// GetProductByName godoc
// @Summary Get product by name
// @Description Get product by name
// @Tags products
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param name path string true "Name of Product"
// @Param currency query string true "Currency"
// @Failure 404 {object} payload.Response
// @Failure 502 {object} payload.Response
// @Success 200 {object} payload.GetProductSuccess
// @Router /products/name/{name} [get]
func (pc *ProductController) GetProductByName(ctx *gin.Context) {

	// // get name from URL path
	// name := ctx.Request.URL.Query()["name"][0]

	// filter parameter id from context
	name := ctx.Params.ByName("name")

	var currency = ctx.DefaultQuery("currency", "USD")

	// call admin service to get user by email
	product, err := pc.productService.FindProductByName(name, currency)
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
		payload.GetProductSuccess{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Get product success",
			Data:    product,
		})

}

// GetProductByCategory godoc
// @Summary Get product by Category
// @Description Get product by Category
// @Tags products
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param category path string true "Category"
// @Param currency query string true "Currency"
// @Failure 404 {object} payload.Response
// @Failure 502 {object} payload.Response
// @Success 200 {object} payload.GetProductSuccess
// @Router /products/category/{category} [get]
func (pc *ProductController) GetProductByCategory(ctx *gin.Context) {

	// // get category ID from URL path
	// category := ctx.Request.URL.Query()["category"][0]

	// filter parameter id from context
	category := ctx.Params.ByName("category")

	var currency = ctx.DefaultQuery("currency", "USD")

	// call admin service to get user by email
	product, err := pc.productService.FindProductByCategory(category, currency)
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
		payload.GetProductSuccess{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Get product success",
			Data:    product,
		})

}

// AddProduct godoc
// @Summary Create a product
// @Description User create a product
// @Tags products
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param product body payload.RequestCreateProduct true "New Product"
// @Failure 404 {object} payload.Response
// @Failure 502 {object} payload.Response
// @Success 200 {object} payload.GetProductSuccess
// @Router /products/ [post]
func (pc *ProductController) AddProduct(ctx *gin.Context) {

	// prepare a post request from ctx
	var reqProduct *payload.RequestCreateProduct

	// from context, bind a new post info to json
	if err := ctx.ShouldBindJSON(&reqProduct); err != nil {
		ctx.JSON(http.StatusBadRequest,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
		return
	}

	// call post service to create the post
	product, err := pc.productService.CreateProduct(reqProduct)
	if err != nil {
		if strings.Contains(err.Error(), "title already exists") {
			ctx.JSON(http.StatusConflict,
				payload.Response{
					Status:  "fail",
					Code:    http.StatusConflict,
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

	ctx.JSON(http.StatusCreated,
		payload.CreateProductSuccess{
			Status:  "success",
			Code:    http.StatusCreated,
			Message: "Create a product success",
			Data:    models.FilteredResponse(product),
		})

}

// UpdateProduct godoc
// @Summary Update a product
// @Description User update product info
// @Tags products
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param product body models.Product true "Update post"
// @Param id path string true "Product ID"
// @Failure 404 {object} payload.Response
// @Failure 502 {object} payload.Response
// @Success 200 {object} payload.Response
// @Router /products/{id} [patch]
func (pc *ProductController) UpdateProduct(ctx *gin.Context) {

	// get post ID from URL path
	id := ctx.Param("id")

	// from context, bind a new post info to json
	var product *models.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadGateway,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusBadGateway,
				Message: err.Error(),
			})
		return
	}

	// call post service to update info
	updatedProduct, err := pc.productService.UpdateProduct(id, product)
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
		payload.UpdateProductSuccess{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Update a exist post success",
			Data:    updatedProduct,
		})
}

// DeleteProduct godoc
// @Summary Delete a post by ID
// @Description User delete the product by product ID
// @Tags products
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id path string true "Product ID"
// @Failure 404 {object} payload.Response
// @Failure 502 {object} payload.Response
// @Success 204 {object} payload.Response
// @Router /products/{id} [delete]
func (pc *ProductController) DeleteProductByID(ctx *gin.Context) {
	// get post ID from URL path
	id := ctx.Param("id")

	// call post service to delete post by ID
	err := pc.productService.DeleteProduct(id)

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
			Message: "Delete product success",
		})
}
