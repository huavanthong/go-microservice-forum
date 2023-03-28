package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/infrastructure/persistence/mongodb"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/infrastructure/redis"
	"github.com/your-username/your-project/domain"
)

func main() {
	// Connect to Redis
	redisClient := redis.NewRedisClient()
	defer redisClient.Close()

	// Connect to MongoDB
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = mongoClient.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer mongoClient.Disconnect(context.Background())

	// Create Redis and MongoDB repositories
	redisRepo := redis.NewRedisBasketRepository(redisClient, context.Background())
	mongoRepo := mongodb.NewMongoDBBasketRepository(mongoClient, "basket", "carts")

	// Create BasketService with Redis and MongoDB repositories
	basketService := domain.NewBasketService(redisRepo, mongoRepo)

	// Test BasketService
	userName := "john.doe"
	// Create initial basket
	basket, err := basketService.CreateBasket(userName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Initial basket for user %s: %+v\n", userName, basket)

	// Add item to basket
	item := domain.BasketItem{Name: "iPhone", Price: 1000}
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
