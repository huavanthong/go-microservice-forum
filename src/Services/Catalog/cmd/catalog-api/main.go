package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"go.uber.org/zap"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/api/handlers"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/api/routers"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/domain/services"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/infrastructure/configs"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/infrastructure/storage/mongodb"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

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

	// Repositories setting
	productRepo  *mongodb.ProductRepository
	categoryRepo *mongodb.CategoryRepository

	// Services setting
	productService  services.CatalogServiceImpl
	categoryService services.CategoryServiceImpl

	// Handler setting
	catalogHandler  handlers.CatalogHandler
	categoryHandler handlers.CategoryHandler
)

func dropCollections(db *mongo.Database) error {

	// err := db.Collection("products").Drop()

	// err = db.Collection("category").Drop()

	// err = db.Collection("subcategory").Drop()

	return nil
}

func initCollections(db *mongo.Database) error {

	/* Initialize data for Category */
	dataCategory, err := ioutil.ReadFile("common/data/categorydata.json")
	if err != nil {
		return err
	}

	var categories []interface{}
	if err := json.Unmarshal(dataCategory, &categories); err != nil {
		return err
	}

	categoriesResult, err := db.Collection("category").InsertMany(ctx, categories)
	fmt.Printf("Inserted %v documents into category collection!\n", categoriesResult)

	/* Initialize data for SubCategory */
	dataSubCategory, err := ioutil.ReadFile("common/data/subcategory.json")
	if err != nil {
		return err
	}

	var subCategories []interface{}
	if err := json.Unmarshal(dataSubCategory, &subCategories); err != nil {
		return err
	}

	subCategoriesResult, err := db.Collection("subcategory").InsertMany(ctx, subCategories)
	fmt.Printf("Inserted %v documents into sub category collection!\n", subCategoriesResult)

	return err
}

func handleFlags(db *mongo.Database) {

	enableDummyData := flag.Bool("enable-data", false, "Enable dummy data for testing")
	initData := flag.Bool("init-data", false, "Set this flag if DB should be initialized with dummy data")
	drop := flag.Bool("drop-table", false, "Set this flag if you wan't to drop all user data in your DB")
	flag.Parse()

	if *enableDummyData {
		if *drop {
			msg := ""
			if err := dropCollections(db); err != nil {
				msg = fmt.Sprintf("Error dropping table: %v", err)
			} else {
				msg = "Dropped all collections in DB"
			}

			fmt.Println(msg)
			os.Exit(0)
		}

		if *initData {
			msg := ""
			if err := initCollections(db); err != nil {
				msg = fmt.Sprintf("Error initializing data in DB: %v", err)
			} else {
				msg = "Initialized data in DB."
			}

			fmt.Println(msg)
			os.Exit(0)
		}
	}

}

func init() {

	// Loading config from variable environment
	config, err := configs.LoadConfig("./internal/infrastructure/configs")
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

	/************************ import data for testing on MongoDB  ************************/
	handleFlags(mongoclient.Database("golang_mongodb"))

	// Initialize MongoDB
	productCollection = mongoclient.Database("golang_mongodb").Collection("products")
	categoryCollection = mongoclient.Database("golang_mongodb").Collection("category")

	// Initialize Repository
	productRepo = mongodb.NewProductRepository(logger, productCollection, ctx)
	categoryRepo = mongodb.NewCategoryRepository(logger, categoryCollection, ctx)

	// Initialize services
	productService := services.NewCatalogServiceImpl(logger, productRepo, ctx)
	categoryService := services.NewCategoryServiceImpl(logger, categoryRepo, ctx)

	// Initialize handlers
	catalogHandler = handlers.NewCatalogHandler(logger, productService)
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

// @host localhost:9090
// @BasePath /api/v3
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

	docs.SwaggerInfo.BasePath = "/api/v3"

	/************************ Server routing  *************************/
	router := server.Group("/api/v3")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Hello World"})
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	log.Println("Starting server on port 9090")
	log.Fatal(server.Run(":" + config.Port))
	// 		go func() {
	// 	}()
}
