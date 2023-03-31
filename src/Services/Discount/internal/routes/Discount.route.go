package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/controllers"
)

func SetupProductRouter(router *gin.Engine, discountController controllers.DiscountController) {
	discountRouter := router.Group("/discount")
	{
		discountRouter.GET("/:productName", discountController.GetDiscount)
		discountRouter.POST("/", discountController.CreateDiscount)
		discountRouter.PUT("/", discountController.UpdateDiscount)
		discountRouter.DELETE("/:productName", discountController.DeleteDiscount)

		// Health check
		discountRouter.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "OK"})
		})
	}
}
