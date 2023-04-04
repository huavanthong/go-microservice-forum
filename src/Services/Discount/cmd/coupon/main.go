package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"github.com/jmoiron/sqlx"

	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/config"
	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/controllers"
	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/proto/discountpb"
	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/repositories"
	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/routes"
	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/services"
	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/utils"
)

// Declare global variable
var (
	configPath string = "./internal/config/config.yml"
	server     *gin.Engine
	db         *sqlx.DB

	userService         services.UserService
	UserController      controllers.UserController
	UserRouteController routes.UserRouteController

	authService         services.AuthService
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController
)

func init() {
	// Load configuration from file
	config, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	// Initialize context in background
	ctx := context.TODO()

}
func main() {

	// Load configuration from file
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Set up logger
	logger, err := utils.NewLogger(cfg.LogLevel)
	if err != nil {
		log.Fatalf("Failed to Initialize logger: %v", err)
	}

	/*****************************************************************/
	// Set up database connection
	discountRepos, err := repositories.NewPostgresDBDiscountRepository(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to Initialize Postgres discount database: %v", err)
	}

	// Create DiscountService with Postgres repositories
	discountService := services.NewBasketService(logger, discountRepos)

	// Set up gRPC server
	grpcServer := grpc.NewServer()
	discountServer := services.NewDiscountService(DiscountRepos, logger)
	discountpb.RegisterDiscountServer(grpcServer, discountServer)

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
