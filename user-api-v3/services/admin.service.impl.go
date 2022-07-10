package services

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type AdminServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewAdminService(collection *mongo.Collection, ctx context.Context) AdminServiceImpl {
	return &AdminServiceImpl{collection, ctx}
}
