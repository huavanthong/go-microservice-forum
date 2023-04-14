package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/controllers"
	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/services"
)

type DiscountRouteController struct {
	discountController controllers.DiscountController
}

func NewRouteDiscountController(discountController controllers.DiscountController) DiscountRouteController {
	return DiscountRouteController{discountController}
}

func (dc *DiscountRouteController) DiscountRoute(rg *gin.RouterGroup, discountService services.DiscountService) {

	router := rg.Group("discounts")
	router.GET("/:id", dc.discountController.GetDiscount)
	router.GET("/", dc.discountController.GetAllDiscounts)
	router.POST("/", dc.discountController.CreateDiscount)
	router.PUT("/", dc.discountController.UpdateDiscount)
	router.DELETE("/:id", dc.discountController.DeleteDiscount)

	// Health check
	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})
}
