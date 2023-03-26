package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/api/models"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/domain/services"
)

type CategoryHandler struct {
	log             *zap.Logger
	categoryService services.CategoryService
}

func NewCategoryHandler(log *zap.Logger, categoryService services.CategoryService) CategoryHandler {
	return CategoryHandler{log, categoryService}
}

// GetAllCategories godoc
// @Summary List all existing categories
// @Description List all existing categories of store
// @Tags category
// @Accept  json
// @Produce  json
// @Param page path string true "Page"
// @Param limit path string true "Limit"
// @Failure 500 {object} models.Response
// @Success 200 {array} entities.Category
// @Router /category [get]
func (cc *CategoryHandler) GetAllCategories(ctx *gin.Context) {

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
			models.Response{
				Status:  "fail",
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusOK,
		models.GetAllCategoriesSuccess{
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
// @Failure 404 {object} models.Response
// @Failure 502 {object} models.Response
// @Success 200 {object} models.GetCategorySuccess
// @Router /category/{id} [get]
func (cc *CategoryHandler) GetCategoryByID(ctx *gin.Context) {

	// filter parameter id from context
	categoryId := ctx.Params.ByName("id")

	// find category by id
	category, err := cc.categoryService.FindCategoryByID(categoryId)

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
		models.GetCategorySuccess{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Get category success",
			Data:    category,
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
// @Failure 404 {object} models.Response
// @Failure 502 {object} models.Response
// @Success 200 {object} models.GetCategorySuccess
// @Router /category/name/{name} [get]
func (cc *CategoryHandler) GetCategoryByName(ctx *gin.Context) {

	// filter parameter id from context
	name := ctx.Params.ByName("name")

	// call category service to get category by name
	categories, err := cc.categoryService.FindCategoryByName(name)
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
		models.GetAllCategoriesSuccess{
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
// @Param category body models.RequestCreateCategory true "New Category"
// @Failure 404 {object} models.Response
// @Failure 502 {object} models.Response
// @Success 200 {object} models.GetCategorySuccess
// @Router /category/ [post]
func (cc *CategoryHandler) AddCategory(ctx *gin.Context) {

	// prepare a post request from ctx
	var reqCategory *models.RequestCreateCategory

	// from context, bind a new post info to json
	if err := ctx.ShouldBindJSON(&reqCategory); err != nil {
		ctx.JSON(http.StatusBadRequest,
			models.Response{
				Status:  "fail",
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
		return
	}

	// call post service to create the post
	category, err := cc.categoryService.CreateCategory(reqCategory)
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
		models.CreateCategorySuccess{
			Status:  "success",
			Code:    http.StatusCreated,
			Message: "Create a new category success",
			Data:    category,
		})

}

// UpdateCategory godoc
// @Summary Update a category
// @Description User update category info
// @Tags category
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param product body models.RequestUpdateProduct true "Update post"
// @Param id path string true "Category ID"
// @Failure 404 {object} models.Response
// @Failure 502 {object} models.Response
// @Success 200 {object} models.Response
// @Router /category/{id} [patch]
func (cc *CategoryHandler) UpdateCategory(ctx *gin.Context) {

	// get post ID from URL path
	id := ctx.Param("id")

	// from context, bind a new post info to json
	var reqUpdateCategory *models.RequestUpdateCategory
	if err := ctx.ShouldBindJSON(&reqUpdateCategory); err != nil {
		ctx.JSON(http.StatusBadGateway,
			models.Response{
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
		models.UpdateCategorySuccess{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Update a exist category success",
			Data:    updatedCategory,
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
// @Failure 404 {object} models.Response
// @Failure 502 {object} models.Response
// @Success 204 {object} models.Response
// @Router /category/{id} [delete]
func (cc *CategoryHandler) DeleteCategoryByID(ctx *gin.Context) {
	// get post ID from URL path
	id := ctx.Param("id")

	// call category service to delete category by ID
	err := cc.categoryService.DeleteCategory(id)

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
			Message: "Delete category success",
		})
}
