package interfaces

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/services"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/infrastructure/config"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/interfaces/api/routes"

	"github.com/sirupsen/logrus"
)

type Server struct {
	config        *config.Config
	logger        *logrus.Logger
	basketService *services.BasketService
}

func NewServer(cfg *config.Config, logger *logrus.Logger, basketService *services.BasketService) (*Server, error) {
	return &Server{
		config:        cfg,
		logger:        logger,
		basketService: basketService,
	}, nil
}

func (s *Server) Start() error {
	// Initialize router engine by Gin
	router := gin.Default()

	// Register routes from package routes
	routes.RegisterRoutes(router, s.basketService)

	port := fmt.Sprintf(":%d", s.config.Port)
	s.logger.Infof("Starting server on port %d", s.config.Port)
	if err := http.ListenAndServe(port, router); err != nil {
		return fmt.Errorf("failed to listen and serve on port %d: %w", s.config.Port, err)
	}

	return nil
}
