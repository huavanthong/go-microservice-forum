package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/api/handlers"
)

func SetupProductRouter(router *gin.RouterGroup, catalogHandler handlers.CatalogHandler) {
	productRouter := router.Group("/products")
	{
		// productRouter.Use(middleware.MiddlewareValidateProduct())
		productRouter.GET("/", catalogHandler.GetAllProducts)
		productRouter.GET("/:id", catalogHandler.GetProductByID)
		productRouter.GET("/name/:name", catalogHandler.GetProductByName)
		productRouter.GET("/category/:category", catalogHandler.GetProductByCategory)
		productRouter.PATCH("/:id", catalogHandler.UpdateProduct)
		productRouter.DELETE("/:id", catalogHandler.DeleteProductByID)
		productRouter.POST("/", catalogHandler.AddProduct)

		// Health check
		productRouter.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "OK"})
		})
	}
}
