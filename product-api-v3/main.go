package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.uber.org/zap"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/huavanthong/microservice-golang/product-api-v3/config"
	"github.com/huavanthong/microservice-golang/product-api-v3/controllers"
	"github.com/huavanthong/microservice-golang/product-api-v3/routes"
	"github.com/huavanthong/microservice-golang/product-api-v3/services"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	docs "github.com/huavanthong/microservice-golang/product-api-v3/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	// Server setting
	server      *gin.Engine     // The framework's instance, it contains the muxer, middleware and configuration settings.
	ctx         context.Context // Context running in background
	mongoclient *mongo.Client   // MongoDB
	logger      *zap.Logger

	// Product Controller setting
	productService         services.ProductService
	ProductController      controllers.ProductController
	productCollection      *mongo.Collection
	ProductRouteController routes.ProductRouteController
)

func init() {

	// Loading config from variable environment
	config, err := config.LoadConfig("./config/")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	// Enable logger
	logger, _ = zap.NewProduction()

	// Init an context running in background
	ctx = context.TODO()

	// Connect to MongoDB
	mongoconn := options.Client().ApplyURI(config.DBUri)
	mongoclient, err := mongo.Connect(ctx, mongoconn)

	if err != nil {
		panic(err)
	}

	if err := mongoclient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("MongoDB successfully connected...")

	// Add the Product Service, Controllers and Routes
	productCollection = mongoclient.Database("golang_mongodb").Collection("products")
	productService = services.NewProductServiceImpl(logger, productCollection, ctx)
	ProductController = controllers.NewProductController(logger, productService)
	ProductRouteController = routes.NewRouteProductController(ProductController)

	// Default returns an Engine instance with the Logger and Recovery middleware already attached.
	server = gin.Default()
}

// @title UserManagement Service API Document
// @version 1.0
// @description List APIs of UserManagement Service
// @termsOfService http://swagger.io/terms/

// @host localhost:8000
// @BasePath /api/v3
func main() {

	/************************ Init MongoDB *************************/
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("Could not load config", err)
	}

	defer mongoclient.Disconnect(ctx)

	/************************ Start internal server *************************/
	startGinServer(config)

}

func startGinServer(config config.Config) {

	/************************ Allow Cross Orgin Resource Sharing  *************************/
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", "http://localhost:3000"}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	docs.SwaggerInfo.BasePath = "/api/v3"

	/************************ Server routing  *************************/
	router := server.Group("/api/v3")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Hello World"})
	})

	/************************ Controller  *************************/
	ProductRouteController.ProductRoute(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	go func() {
		log.Println("Starting server on port 9090")
		log.Fatal(server.Run(":" + config.Port))
	}()
}
