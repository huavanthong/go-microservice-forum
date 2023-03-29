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

	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/services"

	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/infrastructure/config"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/infrastructure/persistence/mongodb"
	redisdb "github.com/huavanthong/microservice-golang/src/Services/Basket/internal/infrastructure/persistence/redis"

	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/interfaces"
)

func main() {

	// Loading config from variable environment
	config, err := config.LoadConfig("./internal/infrastructure/config")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	// Init an context running in background
	ctx := context.TODO()

	// Connect to Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr: config.RedisUri,
	})
	defer redisClient.Close()

	if _, err := redisClient.Ping(ctx).Result(); err != nil {
		panic(err)
	}

	err = redisClient.Set(ctx, "test", "Welcome to Golang with Redis and MongoDB", 0).Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("Redis client connected successfully...")

	// Connect to MongoDB
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(config.DBUri))
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
	// Create Redis and MongoDB repositories
	redisRepo := redisdb.NewRedisBasketRepository(redisClient, context.Background())
	mongoRepo := mongodb.NewMongoDBBasketRepository(mongoClient, "basket", "carts")

	// Create BasketService with Redis and MongoDB repositories
	basketService := services.NewBasketService(redisRepo, mongoRepo)

	// Create a new instance of the logger.
	log := logrus.New()

	// Create a new instance of the server
	srv, err := interfaces.NewServer(config, log, basketService)
	if err != nil {
		log.Fatalf("failed to create server instance: %v", err)
	}

	// Start the server
	go func() {
		if err := srv.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start server: %v", err)
		}
	}()
}
