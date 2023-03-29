package main

import (
	"context"
	"fmt"
	"log"

	redis "github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/entities"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/services"

	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/infrastructure/config"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/infrastructure/persistence/mongodb"
	redisdb "github.com/huavanthong/microservice-golang/src/Services/Basket/internal/infrastructure/persistence/redis"
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

	// Test BasketService
	userName := "john.doe"
	// Create initial basket
	basket, err := basketService.CreateBasket(userName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Initial basket for user %s: %+v\n", userName, basket)

	// Add item to basket
	item := entities.ShoppingCartItem{
		ProductName: "iPhone",
		Price:       1000,
	}
	basket.Items = append(basket.Items, item)
	basket, err = basketService.UpdateBasket(userName, basket)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated basket for user %s: %+v\n", userName, basket)

	// Delete basket
	err = basketService.DeleteBasket(userName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted basket for user %s\n", userName)
}
