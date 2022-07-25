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

func NewAdminServiceImpl(collection *mongo.Collection, ctx context.Context) AdminService {
	return &AdminServiceImpl{collection, ctx}
}

func (ac *AdminServiceImpl) GetAllUsers(page int, limit int) ([]*models.User, error) {

	// page return product
	if page == 0 {
		page = 1
	}

	// limit data return
	if limit == 0 {
		limit = 10
	}

	skip := (page - 1) * limit

	opt := options.FindOptions{}
	opt.SetLimit(int64(limit))
	opt.SetSkip(int64(skip))

	// create a query command
	query := bson.M{}

	// find all posts with optional data
	cursor, err := ac.collection.Find(ac.ctx, query, &opt)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ac.ctx)

	// create container for data
	var users []*models.User

	// with data find out, we will decode them and append to array
	for cursor.Next(ac.ctx) {
		user := &models.User{}
		err := cursor.Decode(user)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	// if any item error, return err
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// if data is empty, return nil
	if len(users) == 0 {
		return []*models.User{}, nil
	}

	return users, nil
}

func (ac *AdminServiceImpl) GetUserByID(string id) (*models.User, error) {

	// convert string id to object id
	oid, _ := primitive.ObjectIDFromHex(id)

	// create a query command
	query := bson.M{"_id": oid}

	// create container for data
	var user *models.User

	// find one user by query command
	err := ac.collection.FindOne(ac.ctx, query).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &models.User{}, err
		}
		return nil, err
	}

	return user, nil
}

func (ac *AdminServiceImpl) GetUserByParam(*models.SignInInput) (*models.DBResponse, error) {
	panic("implement me")
}
