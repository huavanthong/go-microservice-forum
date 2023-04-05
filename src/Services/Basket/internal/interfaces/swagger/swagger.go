package swagger

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	docs "github.com/huavanthong/microservice-golang/src/Services/Basket/docs"
)

func RegisterSwagger(router *gin.RouterGroup) {
	docs.SwaggerInfo.BasePath = "/api/v1"

	// Set up Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
