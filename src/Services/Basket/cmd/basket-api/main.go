package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

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
)

var (
	configPath = "./internal/infrastructure/config"

	server      interfaces.Server
	ctx         context.Context
	mongoclient *mongo.Client
	redisclient *redis.Client

	mongoPersistence persistence.BasketPersistence
	redisPersistence persistence.BasketPersistence
	basketRepository repositories.BasketRepository
	basketService    services.BasketService
)

func GetValue(client *redis.Client, ctx context.Context) error {

	data := client.Get(ctx, "test")
	fmt.Println("Check value from key test: ", data)

	return nil
}

func init() {
	// Loading config from variable environment
	_, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	// Init context running in background
	ctx = context.TODO()

	// Connect to Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer redisClient.Close()

	if _, err := redisClient.Ping(ctx).Result(); err != nil {
		log.Fatal(err)
	}

	err = redisClient.Set(ctx, "test", "Welcome to Golang with Redis and MongoDB", 0).Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Redis client connected successfully...")

	// Connect to MongoDB
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27018"))
	if err != nil {
		log.Fatal(err)
	}
	err = mongoClient.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer mongoClient.Disconnect(context.Background())

	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("MongoDB successfully connected...")

	/*****************************************************************/
	// Create Redis and MongoDB persistence
	mongoPersistence = mongodb.NewMongoDBBasketPersistence(mongoClient, "basket", "carts")
	redisPersistence = redisdb.NewRedisBasketPersistence(redisClient, ctx)

	// Create basket repositories
	basketRepository = repositories.NewBasketRepositoryImpl(mongoPersistence, redisPersistence)

	// Create BasketService with Redis and MongoDB repositories
	basketService = services.NewBasketServiceImpl(basketRepository)

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

	// Create a new instance of the logger.
	log := logrus.New()

	// Create a new instance of the server
	srv, err := interfaces.NewServer(config, log, basketService)
	if err != nil {
		log.Fatalf("failed to create server instance: %v", err)
	}

	// Start the server
	err = srv.Start()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("failed to start server: %v", err)
	}
}
