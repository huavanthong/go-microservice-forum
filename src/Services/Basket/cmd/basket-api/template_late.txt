package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"

	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/services"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/infrastructure/persistence/mongodb"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/infrastructure/persistence/redis"
)

func main() {
	// Connect to Redis
	redisClient := redis.NewRedisClient()
	defer redisClient.Close()

	// Connect to MongoDB
	connectionString := "mongodb://localhost:27017"
	database := "basket"
	mongoClientCustomer, err := mongodb.NewNewMongoDBClient(connectionString, database)
	if err != nil {
		log.Fatal(err)
	}
	defer mongoClientCustomer.Disconnect()

	// Create Redis and MongoDB repositories
	redisRepo := redis.NewRedisBasketRepository(redisClient, context.Background())
	mongoRepo := mongodb.NewMongoDBBasketRepository(mongoClientCustomer, "basket", "carts")

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
	item := services.BasketItem{Name: "iPhone", Price: 1000}
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
