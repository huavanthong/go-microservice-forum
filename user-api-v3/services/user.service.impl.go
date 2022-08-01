package services

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/huavanthong/microservice-golang/user-api-v3/models"
	"github.com/huavanthong/microservice-golang/user-api-v3/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewUserServiceImpl(collection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{collection, ctx}
}

// FindUserByID
func (us *UserServiceImpl) FindUserById(id string) (*models.DBResponse, error) {

	// convert string id to object id
	oid, _ := primitive.ObjectIDFromHex(id)

	// create a query command
	query := bson.M{"_id": oid}

	// create container for data
	var user *models.DBResponse

	// find one user by query command
	err := us.collection.FindOne(us.ctx, query).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &models.DBResponse{}, err
		}
		return nil, err
	}

	return user, nil
}

// FindUserByEmail
func (us *UserServiceImpl) FindUserByEmail(email string) (*models.DBResponse, error) {

	// create container for data
	var user *models.DBResponse

	// create a query command
	query := bson.M{"email": strings.ToLower(email)}

	// find one user by query command
	err := us.collection.FindOne(us.ctx, query).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &models.DBResponse{}, err
		}
		return nil, err
	}

	return user, nil
}

// UpsertUser
func (uc *UserServiceImpl) UpsertUser(email string, data *models.UpdateDBUser) (*models.DBResponse, error) {
	doc, err := utils.ToDoc(data)
	if err != nil {
		return nil, err
	}

	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(1)
	query := bson.D{{Key: "email", Value: email}}
	update := bson.D{{Key: "$set", Value: doc}}
	res := uc.collection.FindOneAndUpdate(uc.ctx, query, update, opts)

	var updatedPost *models.DBResponse

	if err := res.Decode(&updatedPost); err != nil {
		return nil, errors.New("no post with that Id exists")
	}

	return updatedPost, nil
}

func (uc *UserServiceImpl) UpdateUserById(id string, field string, value string) (*models.DBResponse, error) {

	// convert string id to object id
	objUserId, _ := primitive.ObjectIDFromHex(id)

	// create a query command
	query := bson.D{{Key: "_id", Value: objUserId}}

	// create a updated data
	update := bson.D{{Key: "$set", Value: bson.D{{Key: field, Value: value}}}}

	// call mongo driver to update data
	result, err := uc.collection.UpdateOne(uc.ctx, query, update)

	fmt.Print(result.ModifiedCount)

	if err != nil {
		fmt.Print(err)
		return &models.DBResponse{}, err
	}

	return &models.DBResponse{}, nil
}

// func (uc *UserServiceImpl) UpdateOne(field string, value interface{}) (*models.DBResponse, error) {
// 	query := bson.D{{Key: field, Value: value}}
// 	update := bson.D{{Key: "$set", Value: bson.D{{Key: field, Value: value}}}}
// 	result, err := uc.collection.UpdateOne(uc.ctx, query, update)

// 	fmt.Print(result.ModifiedCount)

// 	if err != nil {
// 		fmt.Print(err)
// 		return &models.DBResponse{}, err
// 	}

// 	return &models.DBResponse{}, nil
// }
