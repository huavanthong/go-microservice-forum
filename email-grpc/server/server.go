package server

import (
	pb "github.com/huavanthong/microservice-golang/email-grpc/proto/email"
	"github.com/huavanthong/microservice-golang/user-api-v3/config"
	"github.com/huavanthong/microservice-golang/user-api-v3/services"
	"go.mongodb.org/mongo-driver/mongo"
)

// define email server
type Server struct {
	pb.UnimplementedAuthServiceServer
	config         config.Config
	authService    services.AuthService
	userService    services.UserService
	userCollection *mongo.Collection
}

// Constructor for email grpc server
func NewGrpcServer(config config.Config, authService services.AuthService,
	userService services.UserService, userCollection *mongo.Collection) (*Server, error) {

	server := &Server{
		config:         config,
		authService:    authService,
		userService:    userService,
		userCollection: userCollection,
	}

	return server, nil
}
