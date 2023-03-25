package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/api/handlers"
)

func SetupCategoryRouter(router *gin.Engine, categoryHandler handlers.CategoryHandler) {
	categoryRouter := router.Group("/categories")
	{
		categoryRouter.GET("/", categoryHandler.GetAllCategories)
		categoryRouter.GET("/:id", categoryHandler.GetCategoryByID)
		categoryRouter.GET("/name/:name", categoryHandler.GetCategoryByName)
		categoryRouter.POST("/", categoryHandler.AddCategory)
		categoryRouter.PATCH("/:id", categoryHandler.UpdateCategory)
		categoryRouter.DELETE("/:id", categoryHandler.DeleteCategoryByID)

		// Health check
		categoryRouter.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "OK"})
		})
	}
}
