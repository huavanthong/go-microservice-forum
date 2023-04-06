package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	redis "github.com/go-redis/redis/v8"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/infrastructure/persistence/mongodb"
	redisdb "github.com/huavanthong/microservice-golang/src/Services/Basket/internal/infrastructure/persistence/redis"

	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/repositories"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/services"

	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/interfaces/api/models"
)

func main() {
	// Init context in background
	ctx := context.TODO()

	// Khởi tạo Redis Client
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer redisClient.Close()

	if _, err := redisClient.Ping(ctx).Result(); err != nil {
		panic(err)
	}

	err := redisClient.Set(ctx, "test", "Welcome to Golang with Redis and MongoDB", 0).Err()
	if err != nil {
		panic(err)
	}

	// You can move this to another layer
	data := redisClient.Get(ctx, "test")
	fmt.Println("Check value from key test: ", data)

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

	// Khởi tạo RedisBasketPersistence
	redisPersistence := redisdb.NewRedisBasketPersistence(redisClient, ctx)
	mongoPersistence := mongodb.NewMongoDBBasketPersistence(mongoClient, "basket", "carts")

	// Sử dụng RedisBasketPersistence để thao tác với Redis
	//_, _ = redisPersistence.Create("example-user")

	// Create basket repositories
	basketRepository := repositories.NewBasketRepositoryImpl(mongoPersistence, redisPersistence)

	// Create BasketService with Redis and MongoDB repositories
	basketService := services.NewBasketServiceImpl(basketRepository)

	cbr := models.CreateBasketRequest{
		UserID:   "012345",
		UserName: "hvthong",
	}

	_, _ = basketService.CreateBasket(&cbr)
}
