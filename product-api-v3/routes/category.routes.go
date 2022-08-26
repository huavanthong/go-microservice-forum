package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/huavanthong/microservice-golang/product-api-v3/controllers"
)

type CategoryRouteController struct {
	categoryController controllers.CategoryController
}

func NewRouteCategoryController(categoryController controllers.CategoryController) CategoryRouteController {
	return CategoryRouteController{categoryController}
}

func (cc *CategoryRouteController) CategoryRoute(rg *gin.RouterGroup) {

	router := rg.Group("category")
	router.GET("/", cc.categoryController.GetAllCategories)
	router.GET("/:id", cc.categoryController.GetCategoryByID)
	router.GET("/name/:name", cc.categoryController.GetCategoryByName)
	router.POST("/", cc.categoryController.AddCategory)
	router.PATCH("/:id", cc.categoryController.UpdateCategory)
	router.DELETE("/:id", cc.categoryController.DeleteCategoryByID)
}
