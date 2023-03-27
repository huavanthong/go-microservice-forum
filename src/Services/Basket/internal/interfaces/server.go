package interfaces

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	redis "github.com/go-redis/redis/v8"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/repositories"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/infrastructure/config"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/interfaces/api/controllers"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/interfaces/api/routes"

	"github.com/sirupsen/logrus"
)

type Server struct {
	config           *config.Config
	logger           *logrus.Logger
	redisClient      *redis.Client
	basketController controllers.BasketController
}

func NewServer(cfg *config.Config, logger *logrus.Logger, redisClient *redis.Client, basketController controllers.BasketController) *Server {
	return &Server{
		config:           cfg,
		logger:           logger,
		redisClient:      redisClient,
		basketController: basketController,
	}
}

func (s *Server) Start() error {
	// Khởi tạo router engine bằng Gin
	router := gin.Default()

	// Khởi tạo repositories
	basketRepo := repositories.BasketRepository{}

	// Đăng ký các route từ package router
	routes.RegisterRoutes(router, basketRepo)

	port := fmt.Sprintf(":%d", router.config.Server.Port)
	s.logger.Infof("Starting server on port %d", router.config.Server.Port)
	if err := http.ListenAndServe(port, router); err != nil {
		return fmt.Errorf("failed to listen and serve on port %d: %w", s.config.Server.Port, err)
	}

	return nil
}
