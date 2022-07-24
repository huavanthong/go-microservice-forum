package services

import (
	"context"

	"github.com/huavanthong/microservice-golang/user-api-v3/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type AdminServiceImpl struct {
	adminCollection *mongo.Collection
	ctx             context.Context
}

func NewAdminServiceImpl(collection *mongo.Collection, ctx context.Context) AdminService {
	return &AdminServiceImpl{collection, ctx}
}

func (ac *AdminServiceImpl) GetAllUsers(*models.SignUpInput) ([]*models.DBResponse, error) {

	panic("implement me")

	// query := bson.M{}

	// users, err := ac.adminCollection.Find(ac.ctx, query, nil)
	// if err != nil {
	// 	return nil, err
	// }
	// defer users.Close(ac.ctx)
}

func (ac *AdminServiceImpl) GetUserByID(*models.SignInInput) (*models.DBResponse, error) {
	panic("implement me")
}

func (ac *AdminServiceImpl) GetUserByParam(*models.SignInInput) (*models.DBResponse, error) {
	panic("implement me")
}
