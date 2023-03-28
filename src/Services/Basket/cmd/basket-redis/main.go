package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/infrastructure/config"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/interfaces"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/utils"
)

func main() {
	// Load configurations
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load configurations: %v", err)
	}

	// Connect to Redis
	redisClient := utils.ConnectRedis(config.RedisURL)
	if redisClient == nil {
		log.Fatalf("Failed to connect to Redis")
	}
	defer redisClient.Close()

	// Create a new instance of the server
	srv, err := interfaces.NewServer(cfg)
	if err != nil {
		log.Fatalf("failed to create server instance: %v", err)
	}

	// Start the server
	go func() {
		if err := srv.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	// Handle graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutting down server...")

	// Wait for pending requests to complete before shutting down
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("failed to gracefully shutdown server: %v", err)
	}

	log.Println("server stopped")
}
