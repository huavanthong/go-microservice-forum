package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/huavanthong/microservice-golang/product-api-v3/controllers"
	"github.com/huavanthong/microservice-golang/product-api-v3/middleware"
	"github.com/huavanthong/microservice-golang/product-api-v3/services"
)

type ProductRouteController struct {
	productController controllers.ProductController
}

func NewRouteProductController(productController controllers.ProductController) ProductRouteController {
	return ProductRouteController{productController}
}

func (pc *ProductRouteController) UserRoute(rg *gin.RouterGroup, productService services.ProductService) {

	router := rg.Group("products")
	router.Use(middleware.MiddlewareValidateProduct())
	router.GET("/", pc.productController.GetAllProducts)
	router.GET("/:id", pc.productController.GetProductByID)
	router.GET("/:name", pc.productController.GetProductByName)
	router.GET("/:category", pc.productController.GetProductByCategory)
	router.POST("/", pc.productController.AddProduct)
	router.PATCH("/:id", pc.productController.UpdateProduct)
	router.DELETE("/:id", pc.productController.DeleteProduct)
}
