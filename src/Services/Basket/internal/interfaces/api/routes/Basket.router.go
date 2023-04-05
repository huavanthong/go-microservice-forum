package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/services"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/interfaces/api/controllers"
)

/* Design old: Thiết kế như vậy thì ta cần phải khởi tạo categoryHandler từ bên ngoài */
// func SetupBasketRouter(router *gin.Engine, basketHandler controllers.BasketController) {
// 	basketRoutes := router.Group("/categories")
// 	{
// 		basketRoutes.GET("/:id", basketHandler.GetBasketByID)
// 		basketRoutes.POST("/:id", basketHandler.UpdateBasket)
// 		basketRoutes.DELETE("/:id", basketHandler.DeleteBasket)
// 		basketRoutes.POST("/:id/items", basketHandler.AddItem)
// 		basketRoutes.DELETE("/:id/items/:itemId", basketHandler.DeleteItem)

// 		// Health check
// 		basketRoutes.GET("/health", func(c *gin.Context) {
// 			c.JSON(http.StatusOK, gin.H{"status": "OK"})
// 		})
// 	}
// }

/* Design new: Thiết kế như vậy thì router sẽ bị phụ thuộc vào controller ở trong*/
func RegisterRoutes(router *gin.RouterGroup, basketService services.BasketService) {

	basketController := controllers.NewBasketController(basketService)

	basketRoutes := router.Group("/basket")
	{
		basketRoutes.GET("/:id", basketController.GetBasket)
		basketRoutes.POST("/", basketController.CreateBasket)
		basketRoutes.DELETE("/:id", basketController.DeleteBasket)

		basketRoutes.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "OK"})
		})
	}
}
