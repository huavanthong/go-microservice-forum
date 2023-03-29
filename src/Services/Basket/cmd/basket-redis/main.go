package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/infrastructure/config"

	redis "github.com/go-redis/redis/v8"
)

func main() {
	// Load configurations
	cfg, err := config.LoadConfig("./internal/infrastructure/config")
	if err != nil {
		log.Fatalf("failed to load configurations: %v", err)
	}

	// Init context in background
	ctx := context.TODO()

	// Connect to Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr: cfg.RedisUri,
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

	// // Create a new instance of the server
	// srv, err := interfaces.NewServer(cfg)
	// if err != nil {
	// 	log.Fatalf("failed to create server instance: %v", err)
	// }

	// // Start the server
	// go func() {
	// 	if err := srv.Start(); err != nil && err != http.ErrServerClosed {
	// 		log.Fatalf("failed to start server: %v", err)
	// 	}
	// }()

	// Handle graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutting down server...")

	// Wait for pending requests to complete before shutting down
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// if err := srv.Shutdown(ctx); err != nil {
	// 	log.Fatalf("failed to gracefully shutdown server: %v", err)
	// }

	log.Println("server stopped")
}
