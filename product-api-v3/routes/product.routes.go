package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/huavanthong/microservice-golang/product-api-v3/controllers"
)

type ProductRouteController struct {
	productController controllers.ProductController
}

func NewRouteProductController(productController controllers.ProductController) ProductRouteController {
	return ProductRouteController{productController}
}

func (pc *ProductRouteController) ProductRoute(rg *gin.RouterGroup) {

	router := rg.Group("products")
	// router.Use(middleware.MiddlewareValidateProduct())
	router.GET("/", pc.productController.GetAllProducts)
	router.GET("/:id", pc.productController.GetProductByID)
	router.GET("/name/:ss", pc.productController.GetProductByName)
	router.GET("/category/:category", pc.productController.GetProductByCategory)
	router.POST("/", pc.productController.AddProduct)
	router.PATCH("/:id", pc.productController.UpdateProduct)
	router.DELETE("/:id", pc.productController.DeleteProductByID)
}
