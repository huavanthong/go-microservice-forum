package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	redis "github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/repositories"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/services"

	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/interfaces/persistence"

	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/infrastructure/config"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/infrastructure/persistence/mongodb"
	redisdb "github.com/huavanthong/microservice-golang/src/Services/Basket/internal/infrastructure/persistence/redis"

	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/interfaces"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/interfaces/api/routes"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/interfaces/swagger"
)

var (
	configPath       = "./internal/infrastructure/config"
	basketDatabase   = "basket-microservice"
	basketCollection = "basket"

	logger *logrus.Logger
	entry  *logrus.Entry

	server      *gin.Engine // The framework's instance, it contains the muxer, middleware and configuration settings.
	myServer    interfaces.Server
	ctx         context.Context // Context running in background
	mongoClient *mongo.Client   // MongoDB
	redisClient *redis.Client   // For in-memory data store

	mongoPersistence persistence.BasketPersistence
	redisPersistence persistence.BasketPersistence
	basketRepository repositories.BasketRepository
	BasketService    services.BasketService
)

func init() {
	// Loading config from variable environment
	config, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	// Create a new logger
	//logger := logrus.New().WithField("app", "myApp")

	// Khởi tạo logger
	logger = logrus.New()

	// Thiết lập định dạng log
	logger.SetFormatter(&logrus.TextFormatter{})

	// Đặt level cho logger
	logger.Level = logrus.DebugLevel

	// Thiết lập đường dẫn cho file log
	file, err := os.OpenFile("log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.SetOutput(file)
	} else {
		logger.Info("Failed to log to file, using default stderr")
	}

	// Tạo một instance của logrus.Entry từ logger
	entry = logger.WithFields(logrus.Fields{
		"app": "myapp",
		"env": "prod",
	})

	// Sử dụng logger
	entry.Info("Hello, world!")

	// Init context running in background
	ctx = context.TODO()

	// Connect to MongoDB
	mongoClient, err = mongo.NewClient(options.Client().ApplyURI(config.DBUri))
	if err != nil {
		log.Fatal(err)
	}
	err = mongoClient.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("MongoDB successfully connected...")

	// Connect to Redis
	redisClient = redis.NewClient(&redis.Options{
		Addr:     config.RedisUri,
		Password: "",
		DB:       0,
	})
	// Error: happen if you close redis connection, please don't close if you want to use out scope
	// defer redisClient.Close()

	if _, err := redisClient.Ping(ctx).Result(); err != nil {
		log.Fatal(err)
	}

	err = redisClient.Set(ctx, "test", "Welcome to Golang with Redis and MongoDB", 0).Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Redis client connected successfully...")
	/*****************************************************************/
	// Create Redis and MongoDB persistence
	mongoPersistence = mongodb.NewMongoDBBasketPersistence(mongoClient, basketDatabase, basketCollection)
	redisPersistence = redisdb.NewRedisBasketPersistence(entry, redisClient, ctx)

	// Create basket repositories
	basketRepository = repositories.NewBasketRepositoryImpl(entry, mongoPersistence, redisPersistence)

	// Create BasketService with Redis and MongoDB repositories
	BasketService = services.NewBasketServiceImpl(basketRepository)

	// Initialize server engine by Gin
	server = gin.Default()
}

// @title Basket Service API Document
// @version 1.0
// @description List APIs of Basket Service
// @termsOfService http://swagger.io/terms/

// @host localhost:8001
// @BasePath /api/v1
func main() {

	// Loading config from variable environment
	config, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	startGinServer(config)

}

func startGinServer(config config.Config) error {

	/************************ Connect Redis *************************/
	value, err := redisClient.Get(ctx, "test").Result()
	if err == redis.Nil {
		fmt.Println("[Main] key: test does not exist")
	} else if err != nil {
		panic(err)
	}

	// Health check server
	router := server.Group("/api/v1")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": value})
	})

	// Register routes from package routes
	routes.RegisterRoutes(router, BasketService)

	// Register Swagger
	swagger.RegisterSwagger(router)

	fmt.Println("Starting server on port:", config.Port)

	log.Fatal(server.Run(":" + config.Port))

	return nil
}

func DetachServer(config *config.Config) {
	// Create a new instance of the logger.
	log := logrus.New()

	// Create a new instance of the server
	srv, err := interfaces.NewServer(config, log, BasketService)
	if err != nil {
		log.Fatalf("failed to create server instance: %v", err)
	}

	// Start the server
	err = srv.Start()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("failed to start server: %v", err)
	}
}
