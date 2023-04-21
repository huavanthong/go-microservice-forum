package main

import (
	"context"
	"fmt"

	redis "github.com/go-redis/redis/v8"
	redisdb "github.com/huavanthong/microservice-golang/src/Services/Basket/internal/infrastructure/persistence/redis"
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

	// Khởi tạo RedisBasketPersistence
	redisPersistence := redisdb.NewRedisBasketPersistence(redisClient, ctx)

	// Sử dụng RedisBasketPersistence để thao tác với Redis
	_, _ = redisPersistence.Create("example-user")

}
