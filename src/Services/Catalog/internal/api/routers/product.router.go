package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/api/handlers"
)

func SetupProductRouter(router *gin.Engine, productHandler handlers.ProductHandler) {
	productRouter := router.Group("/products")
	{
		// productRouter.Use(middleware.MiddlewareValidateProduct())
		productRouter.GET("/", productHandler.GetAllProducts)
		productRouter.GET("/:id", productHandler.GetProductByID)
		productRouter.GET("/name/:name", productHandler.GetProductByName)
		productRouter.GET("/category/:category", productHandler.GetProductByCategory)
		productRouter.PATCH("/:id", productHandler.UpdateProduct)
		productRouter.DELETE("/:id", productHandler.DeleteProductByID)
		productRouter.POST("/", productHandler.AddProduct)

		// Health check
		productRouter.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "OK"})
		})
	}
}
