package main

import (
	"fmt"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/config"
	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/controllers"
	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/proto/discount"
	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/repositories"
	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/routes"
	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/services"
	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/utils"
)

// Declare global variable
var (
	configPath string = "./internal/config/config.yml"
)

func main() {

	// Load configuration from file
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Set up database connection
	db, err := repositories.NewPostgresDB(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to set up database: %v", err)
	}
	defer db.Close()

	// Set up logger
	logger := utils.NewLogger(cfg.LogLevel)

	// Set up gRPC server
	grpcServer := grpc.NewServer()
	discountServer := services.NewDiscountService(db, logger)
	discount.RegisterDiscountServer(grpcServer, discountServer)

	// Set up HTTP server
	router := gin.Default()
	discountController := controllers.NewDiscountController(discountServer)
	routes.SetupDiscountRoutes(router, discountController)

	// Start gRPC server
	go func() {
		addr := fmt.Sprintf(":%d", cfg.GRPC.Port)
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatalf("Failed to listen on port %d: %v", cfg.GRPC.Port, err)
		}
		log.Printf("gRPC server listening on port %d", cfg.GRPC.Port)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()

	// Start HTTP server
	addr := fmt.Sprintf(":%d", cfg.HTTP.Port)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to serve HTTP server: %v", err)
	}
}
