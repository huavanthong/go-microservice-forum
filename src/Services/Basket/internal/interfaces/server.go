package interfaces

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/infrastructures"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/infrastructures/config"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/interface/controllers"
)

type Server struct {
	config           *config.Config
	logger           *logrus.Logger
	redisClient      infrastructures.RedisClient
	basketController controllers.BasketHandler
}

func NewServer(cfg *config.Config, logger *logrus.Logger, redisClient infrastructures.RedisClient, basketHandler BasketHandler) *Server {
	return &Server{
		config:           cfg,
		logger:           logger,
		redisClient:      redisClient,
		basketController: basketHandler,
	}
}

func (s *Server) Start() error {
	// Khởi tạo router engine bằng Gin
	router := gin.Default()

	// Đăng ký các route từ package router
	router.RegisterRoutes(router)

	port := fmt.Sprintf(":%d", router.config.Server.Port)
	s.logger.Infof("Starting server on port %d", router.config.Server.Port)
	if err := http.ListenAndServe(port, router); err != nil {
		return fmt.Errorf("failed to listen and serve on port %d: %w", s.config.Server.Port, err)
	}

	return nil
}
