package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/api/models"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/domain/entities"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/domain/services"
)

type CatalogHandler struct {
	log            *zap.Logger
	productService services.ProductService
}

func NewCatalogHandler(log *zap.Logger, productService services.ProductService) CatalogHandler {
	return CatalogHandler{log, productService}
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
// @Failure 500 {object} models.Response
// @Success 200 {array} entities.Product
// @Router /products [get]
func (pc *CatalogHandler) GetAllProducts(ctx *gin.Context) {

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
			models.Response{
				Status:  "fail",
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusOK,
		models.ResponseSuccess{
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
// @Failure 404 {object} models.Response
// @Failure 502 {object} models.Response
// @Success 200 {object} models.GetProductSuccess
// @Router /products/{id} [get]
// GetUserByID get a user by id in DB
func (pc *CatalogHandler) GetProductByID(ctx *gin.Context) {

	// filter parameter id from context
	productId := ctx.Params.ByName("id")

	var currency = ctx.DefaultQuery("currency", "USD")

	// find user by id
	product, err := pc.productService.FindProductByID(productId, currency)

	// catch error
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			ctx.JSON(http.StatusNotFound,
				models.Response{
					Status:  "fail",
					Code:    http.StatusNotFound,
					Message: err.Error(),
				})
			return
		}
		ctx.JSON(http.StatusBadGateway,
			models.Response{
				Status:  "fail",
				Code:    http.StatusBadGateway,
				Message: err.Error(),
			})
		return
	}

	// succes
	ctx.JSON(http.StatusOK,
		models.GetProductSuccess{
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
// @Failure 404 {object} models.Response
// @Failure 502 {object} models.Response
// @Success 200 {object} models.GetProductSuccess
// @Router /products/name/{name} [get]
func (pc *CatalogHandler) GetProductByName(ctx *gin.Context) {

	// // get name from URL path
	// name := ctx.Request.URL.Query()["name"][0]

	// filter parameter id from context
	name := ctx.Params.ByName("name")

	var currency = ctx.DefaultQuery("currency", "USD")

	// call admin service to get user by email
	products, err := pc.productService.FindProductByName(name, currency)
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			ctx.JSON(http.StatusNotFound,
				models.Response{
					Status:  "fail",
					Code:    http.StatusNotFound,
					Message: err.Error(),
				})
			return
		}
		ctx.JSON(http.StatusBadGateway,
			models.Response{
				Status:  "fail",
				Code:    http.StatusBadGateway,
				Message: err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusOK,
		models.GetProductsSuccess{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Get product success",
			Data:    products,
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
// @Failure 404 {object} models.Response
// @Failure 502 {object} models.Response
// @Success 200 {object} models.GetProductSuccess
// @Router /products/category/{category} [get]
func (pc *CatalogHandler) GetProductByCategory(ctx *gin.Context) {

	// // get category ID from URL path
	// category := ctx.Request.URL.Query()["category"][0]

	// filter parameter id from context
	category := ctx.Params.ByName("category")

	var currency = ctx.DefaultQuery("currency", "USD")

	// call admin service to get user by email
	products, err := pc.productService.FindProductByCategory(category, currency)
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			ctx.JSON(http.StatusNotFound,
				models.Response{
					Status:  "fail",
					Code:    http.StatusNotFound,
					Message: err.Error(),
				})
			return
		}
		ctx.JSON(http.StatusBadGateway,
			models.Response{
				Status:  "fail",
				Code:    http.StatusBadGateway,
				Message: err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusOK,
		models.GetProductsSuccess{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Get product success",
			Data:    products,
		})

}

func HandlerProduct(h gin.HandlerFunc, decors ...func(gin.HandlerFunc) gin.HandlerFunc) gin.HandlerFunc {
	for i := range decors {
		d := decors[len(decors)-1-i] // iterate in reverse
		h = d(h)
	}
	return h
}

// AddProduct godoc
// @Summary Create a product
// @Description User create a product
// @Tags products
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param product body models.RequestCreateProduct true "New Product"
// @Failure 404 {object} models.Response
// @Failure 502 {object} models.Response
// @Success 200 {object} models.GetProductSuccess
// @Router /products/{productType} [post]
func (pc *CatalogHandler) AddProduct(ctx *gin.Context) {

	// prepare a post request from ctx
	var reqProduct *models.RequestCreateProduct

	// from context, bind a new post info to json
	if err := ctx.ShouldBindJSON(&reqProduct); err != nil {
		ctx.JSON(http.StatusBadRequest,
			models.Response{
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
				models.Response{
					Status:  "fail",
					Code:    http.StatusConflict,
					Message: err.Error(),
				})
			return
		}

		ctx.JSON(http.StatusBadGateway,
			models.Response{
				Status:  "fail",
				Code:    http.StatusBadGateway,
				Message: err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusCreated,
		models.ResponseSuccess{
			Status:  "success",
			Code:    http.StatusCreated,
			Message: "Create a product success",
			Data:    entities.FilteredResponse(product),
		})

}

// UpdateProduct godoc
// @Summary Update a product
// @Description User update product info
// @Tags products
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param product body models.RequestUpdateProduct true "Update post"
// @Param id path string true "Product ID"
// @Failure 404 {object} models.Response
// @Failure 502 {object} models.Response
// @Success 200 {object} models.Response
// @Router /products/{id} [patch]
func (pc *CatalogHandler) UpdateProduct(ctx *gin.Context) {

	// get post ID from URL path
	id := ctx.Param("id")

	// from context, bind a new post info to json
	var reqUpdateProduct *models.RequestUpdateProduct
	if err := ctx.ShouldBindJSON(&reqUpdateProduct); err != nil {
		ctx.JSON(http.StatusBadGateway,
			models.Response{
				Status:  "fail",
				Code:    http.StatusBadGateway,
				Message: err.Error(),
			})
		return
	}

	// call product service to update info
	updatedProduct, err := pc.productService.UpdateProduct(id, reqUpdateProduct)
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			ctx.JSON(http.StatusNotFound,
				models.Response{
					Status:  "fail",
					Code:    http.StatusNotFound,
					Message: err.Error(),
				})
			return
		}
		ctx.JSON(http.StatusBadGateway,
			models.Response{
				Status:  "fail",
				Code:    http.StatusBadGateway,
				Message: err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusOK,
		models.ResponseSuccess{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Update a exist post success",
			Data:    entities.FilteredResponse(updatedProduct),
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
// @Failure 404 {object} models.Response
// @Failure 502 {object} models.Response
// @Success 204 {object} models.Response
// @Router /products/{id} [delete]
func (pc *CatalogHandler) DeleteProductByID(ctx *gin.Context) {
	// get post ID from URL path
	id := ctx.Param("id")

	// call product service to delete product by ID
	err := pc.productService.DeleteProduct(id)

	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			ctx.JSON(http.StatusNotFound,
				models.Response{
					Status:  "fail",
					Code:    http.StatusNotFound,
					Message: err.Error(),
				})
			return
		}
		ctx.JSON(http.StatusBadGateway,
			models.Response{
				Status:  "fail",
				Code:    http.StatusBadGateway,
				Message: err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusNoContent,
		models.Response{
			Status:  "success",
			Code:    http.StatusNoContent,
			Message: "Delete product success",
		})
}
