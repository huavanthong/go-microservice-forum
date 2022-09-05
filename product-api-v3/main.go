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

	// Category Controller setting
	categoryService         services.CategoryService
	CategoryController      controllers.CategoryController
	categoryCollection      *mongo.Collection
	CategoryRouteController routes.CategoryRouteController
)

func dropCollections(db *mongo.Database) error {

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
	config, err := config.LoadConfig("./")
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

	// Add the Product Service, Controllers and Routes
	productCollection = mongoclient.Database("golang_mongodb").Collection("products")
	productService = services.NewProductServiceImpl(logger, productCollection, ctx)
	ProductController = controllers.NewProductController(logger, productService)
	ProductRouteController = routes.NewRouteProductController(ProductController)

	// Add the Category Service, Controllers and Routes
	categoryCollection = mongoclient.Database("golang_mongodb").Collection("category")
	categoryService = services.NewCategoryServiceImpl(logger, categoryCollection, ctx)
	CategoryController = controllers.NewCategoryController(logger, categoryService)
	CategoryRouteController = routes.NewRouteCategoryController(CategoryController)

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
	corsConfig.AllowOrigins = []string{"http://localhost:9090", "http://localhost:3000"}
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
	CategoryRouteController.CategoryRoute(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	log.Println("Starting server on port 9090")
	log.Fatal(server.Run(":" + config.Port))
	// 		go func() {
	// 	}()
}
