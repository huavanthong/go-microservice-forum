package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/jmoiron/sqlx"

	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/config"
	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/controllers"
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

	discountRepository      repositories.DiscountRepository
	discountService         services.DiscountService
	DiscountController      controllers.DiscountController
	DiscountRouteController routes.DiscountRouteController
)

func init() {
	// Load configuration from file
	config, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Create connection string from config
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password, config.Database.DBName)

	// Open connection on PostgreSQL
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// Ping database to ensure connection is valid
	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("PostgreSQL successfully connected...")

	// Set up logger
	logger, err := utils.NewLogger(config.Logger)
	if err != nil {
		log.Fatalf("Failed to Initialize logger: %v", err)
	}

	// Create PostgreSQL instance repositories
	discountRepository = repositories.NewPostgresDBDiscountRepository(db)

	// Create DiscountService with Postgres repositories
	discountService = services.NewDiscountService(logger, discountRepository)

	// Create DiscountController
	DiscountController = controllers.NewDiscountController(discountService)

	// Create DiscountRoute
	DiscountRouteController = routes.NewRouteDiscountController(DiscountController)

	server = gin.Default()

}

func main() {

	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("Could not load config", err)
	}

	// startGinServer(config)
	startGrpcServer(*config)
}

func startGrpcServer(config config.Config) {
	// // Set up gRPC server
	// grpcServer := grpc.NewServer()
	// discountServer := services.NewDiscountService(DiscountRepos, logger)
	// discountpb.RegisterDiscountServer(grpcServer, discountServer)

	// // Set up HTTP server
	// router := gin.Default()
	// discountController := controllers.NewDiscountController(discountServer)
	// routes.SetupDiscountRoutes(router, discountController)

	// // Start gRPC server
	// go func() {
	// 	addr := fmt.Sprintf(":%d", cfg.GRPC.Port)
	// 	lis, err := net.Listen("tcp", addr)
	// 	if err != nil {
	// 		log.Fatalf("Failed to listen on port %d: %v", cfg.GRPC.Port, err)
	// 	}
	// 	log.Printf("gRPC server listening on port %d", cfg.GRPC.Port)
	// 	if err := grpcServer.Serve(lis); err != nil {
	// 		log.Fatalf("Failed to serve gRPC server: %v", err)
	// 	}
	// }()

	// server, err := gapi.NewGrpcServer(config, authService, userService, authCollection)
	// if err != nil {
	// 	log.Fatal("cannot create grpc server: ", err)
	// }

	// grpcServer := grpc.NewServer()
	// pb.RegisterAuthServiceServer(grpcServer, server)
	// reflection.Register(grpcServer)

	// listener, err := net.Listen("tcp", config.GrpcServerAddress)
	// if err != nil {
	// 	log.Fatal("cannot create grpc server: ", err)
	// }

	// log.Printf("start gRPC server on %s", listener.Addr().String())
	// err = grpcServer.Serve(listener)
	// if err != nil {
	// 	log.Fatal("cannot create grpc server: ", err)
	// }
}

func startGinServer(config config.Config) {

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success"})
	})

	DiscountRouteController.DiscountRoute(router, discountService)
	log.Fatal(server.Run(":" + config.App.Port))
}
