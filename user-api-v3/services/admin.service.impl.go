package services

import (
	"context"

	"github.com/huavanthong/microservice-golang/user-api-v3/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type AdminServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewAdminService(collection *mongo.Collection, ctx context.Context) AdminService {
	return &AdminServiceImpl{collection, ctx}
}

func (ac *AdminServiceImpl) GetAllUsers(*models.SignUpInput) (*models.DBResponse, error) {
	panic("implement me")
}

func (ac *AdminServiceImpl) GetUserByID(*models.SignInInput) (*models.DBResponse, error) {
	panic("implement me")
}

func (ac *AdminServiceImpl) GetUserByParam(*models.SignInInput) (*models.DBResponse, error) {
	panic("implement me")
}
