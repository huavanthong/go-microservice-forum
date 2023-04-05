package interfaces

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/services"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/interfaces/api/routes"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/interfaces/swagger"

	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/infrastructure/config"

	"github.com/sirupsen/logrus"
)

type Server struct {
	config        *config.Config
	logger        *logrus.Logger
	basketService services.BasketService
}

func NewServer(cfg *config.Config, logger *logrus.Logger, basketService services.BasketService) (*Server, error) {

	return &Server{
		config:        cfg,
		logger:        logger,
		basketService: basketService,
	}, nil
}

func (s *Server) Start() error {

	// Initialize server engine by Gin
	server := gin.Default()

	// Health check server
	router := server.Group("/api/v1")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Hello World"})
	})

	// Register routes from package routes
	routes.RegisterRoutes(router, s.basketService)

	// Register Swagger
	swagger.RegisterSwagger(router)

	s.logger.Infof("Starting server on port", s.config.Port)

	err := server.Run(":" + s.config.Port)
	if err != nil {
		return fmt.Errorf("failed to listen and serve on port", s.config.Port, err)
	}

	return nil
}
