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

type CategoryController struct {
	log             *zap.Logger
	categoryService services.CategoryService
}

func NewCategoryController(log *zap.Logger, categoryService services.CategoryService) CategoryController {
	return CategoryController{log, categoryService}
}

// GetAllCategories godoc
// @Summary List all existing categories
// @Description List all existing categories of store
// @Tags category
// @Accept  json
// @Produce  json
// @Param page path string true "Page"
// @Param limit path string true "Limit"
// @Failure 500 {object} payload.Response
// @Success 200 {array} models.Category
// @Router /category [get]
func (cc *CategoryController) GetAllCategories(ctx *gin.Context) {

	// get parameter from client
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

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
	categories, err := cc.categoryService.FindAllCategories(intPage, intLimit)
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
			Message: "Get all categories success",
			Data:    categories,
		})
}

// GetCategoryByID godoc
// @Summary Get an category by ID
// @Description Get an category by ID
// @Tags category
// @Accept  json
// @Produce  json
// @Param id path string true "Category ID"
// @Failure 404 {object} payload.Response
// @Failure 502 {object} payload.Response
// @Success 200 {object} payload.GetCategorySuccess
// @Router /category/{id} [get]
func (cc *CategoryController) GetCategoryByID(ctx *gin.Context) {

	// filter parameter id from context
	productId := ctx.Params.ByName("id")

	// find category by id
	product, err := cc.categoryService.FindCategorytByID(productId)

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
			Message: "Get category success",
			Data:    product,
		})

}

// GetCategoryByName godoc
// @Summary Get category by name
// @Description Get category by name
// @Tags category
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param name path string true "Name of Category"
// @Failure 404 {object} payload.Response
// @Failure 502 {object} payload.Response
// @Success 200 {object} payload.GetProductSuccess
// @Router /category/name/{name} [get]
func (cc *CategoryController) GetCategoryByName(ctx *gin.Context) {

	// filter parameter id from context
	name := ctx.Params.ByName("name")

	// call category service to get category by name
	categories, err := cc.categoryService.FindCategoryByName(name)
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
		payload.GetCategorySuccess{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Get category success",
			Data:    categories,
		})

}

// AddCategory godoc
// @Summary Create a new category
// @Description Admin create a new category to storage
// @Tags category
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param product body payload.RequestCreateCategory true "New Category"
// @Failure 404 {object} payload.Response
// @Failure 502 {object} payload.Response
// @Success 200 {object} payload.GetProductSuccess
// @Router /category/ [post]
func (cc *CategoryController) AddCategory(ctx *gin.Context) {

	// prepare a post request from ctx
	var reqCategory *payload.RequestCreateCategory

	// from context, bind a new post info to json
	if err := ctx.ShouldBindJSON(&reqCategory); err != nil {
		ctx.JSON(http.StatusBadRequest,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
		return
	}

	// call post service to create the post
	category, err := cc.categoryService.CreateProduct(reqCategory)
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
		payload.CreateCategorySuccess{
			Status:  "success",
			Code:    http.StatusCreated,
			Message: "Create a new category success",
			Data:    models.FilteredResponse(category),
		})

}

// UpdateCategory godoc
// @Summary Update a category
// @Description User update category info
// @Tags category
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param product body payload.RequestUpdateProduct true "Update post"
// @Param id path string true "Category ID"
// @Failure 404 {object} payload.Response
// @Failure 502 {object} payload.Response
// @Success 200 {object} payload.Response
// @Router /category/{id} [patch]
func (cc *CategoryController) UpdateCategory(ctx *gin.Context) {

	// get post ID from URL path
	id := ctx.Param("id")

	// from context, bind a new post info to json
	var reqUpdateCategory *payload.RequestUpdateCategory
	if err := ctx.ShouldBindJSON(&reqUpdateCategory); err != nil {
		ctx.JSON(http.StatusBadGateway,
			payload.Response{
				Status:  "fail",
				Code:    http.StatusBadGateway,
				Message: err.Error(),
			})
		return
	}

	// call product service to update info
	updatedCategory, err := cc.categoryService.UpdateCategory(id, reqUpdateCategory)
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
		payload.UpdateCategorySuccess{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Update a exist category success",
			Data:    models.FilteredResponse(updatedCategory),
		})
}

// DeleteCategoryByID godoc
// @Summary Delete a category by ID
// @Description User delete the category by category ID
// @Tags category
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id path string true "Product ID"
// @Failure 404 {object} payload.Response
// @Failure 502 {object} payload.Response
// @Success 204 {object} payload.Response
// @Router /category/{id} [delete]
func (cc *CategoryController) DeleteCategoryByID(ctx *gin.Context) {
	// get post ID from URL path
	id := ctx.Param("id")

	// call category service to delete category by ID
	err := cc.categoryService.DeleteCategory(id)

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
			Message: "Delete category success",
		})
}
