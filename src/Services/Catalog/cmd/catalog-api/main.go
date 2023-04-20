package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.uber.org/zap"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/migrations"

	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/api/handlers"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/api/routers"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/domain/repositories"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/domain/services"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/infrastructure/configs"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/infrastructure/storage/mongodb"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	docs "github.com/huavanthong/microservice-golang/src/Services/Catalog/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	// Server setting
	server      *gin.Engine     // The framework's instance, it contains the muxer, middleware and configuration settings.
	ctx         context.Context // Context running in background
	mongoclient *mongo.Client   // MongoDB
	logger      *zap.Logger

	// MongoDB setting
	productCollection  *mongo.Collection
	categoryCollection *mongo.Collection

	// Storage setting
	productStorage  *mongodb.ProductStorage
	categoryStorage *mongodb.CategoryStorage

	// Repositories setting
	productRepo       repositories.ProductRepository
	productSearchRepo repositories.ProductSearchRepository
	categoryRepo      repositories.CategoryRepository

	// Handler setting
	catalogHandler  handlers.CatalogHandler
	categoryHandler handlers.CategoryHandler

	// Services setting
	catalogtService services.CatalogService
	categoryService services.CategoryService
)

func init() {

	// Loading config from variable environment
	cfg, err := configs.LoadConfig("./internal/infrastructure/configs")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	// Enable logger
	logger, _ = zap.NewProduction()

	// Init an context running in background
	ctx = context.TODO()

	// Connect to MongoDB
	mongoconn := options.Client().ApplyURI(cfg.DBContainerUri)
	mongoclient, err := mongo.Connect(ctx, mongoconn)

	if err != nil {
		panic(err)
	}

	if err := mongoclient.Ping(ctx, nil); err != nil {
		panic(err)
	}

	fmt.Println("MongoDB successfully connected...")

	/************************ import data for testing on MongoDB  ************************/
	migrations.Migrations(mongoclient.Database(cfg.DBName), ctx)

	// Initialize MongoDB
	productCollection = mongoclient.Database(cfg.DBName).Collection(cfg.DBCollProduct)
	categoryCollection = mongoclient.Database(cfg.DBName).Collection(cfg.DBCollCategory)

	// Initialize Storage
	productStorage = mongodb.NewProductStorage(logger, productCollection, ctx)
	categoryStorage = mongodb.NewCategoryStorage(logger, categoryCollection, ctx)

	// Initialize Repository
	productRepo = repositories.NewProductRepositoryImpl(productStorage)
	productSearchRepo = repositories.NewProductSearchRepositoryImpl(productStorage)
	categoryRepo = repositories.NewCategoryRepositoryImpl(categoryStorage)

	// Initialize services
	catalogtService := services.NewCatalogServiceImpl(logger, productRepo, productSearchRepo, ctx)
	categoryService := services.NewCategoryServiceImpl(logger, categoryRepo, ctx)

	// Initialize handlers
	catalogHandler = handlers.NewCatalogHandler(logger, catalogtService)
	categoryHandler = handlers.NewCategoryHandler(logger, categoryService)

	// Initialize middleware
	//authMiddleware := middleware.NewAuthMiddleware()

	// Initialize router
	router := gin.Default()

	// Setup routes
	routers.SetupCategoryRouter(router, categoryHandler)
	routers.SetupProductRouter(router, catalogHandler)

	// Default returns an Engine instance with the Logger and Recovery middleware already attached.
	server = gin.Default()
}

// @title UserManagement Service API Document
// @version 1.0
// @description List APIs of UserManagement Service
// @termsOfService http://swagger.io/terms/

// @host localhost:8000
// @BasePath /api/v1
func main() {

	/************************ Init MongoDB *************************/
	config, err := configs.LoadConfig(".")

	if err != nil {
		log.Fatal("Could not load config", err)
	}

	defer mongoclient.Disconnect(ctx)

	/************************ Start internal server *************************/
	startGinServer(config)

}

func startGinServer(config configs.Config) {

	/************************ Allow Cross Orgin Resource Sharing  *************************/
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:9090", "http://localhost:3000"}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	docs.SwaggerInfo.BasePath = "/api/v1"

	/************************ Server routing  *************************/
	router := server.Group("/api/v1")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Hello World"})
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	log.Println("Starting server on port", config.Port)
	log.Fatal(server.Run(":" + config.Port))
}
