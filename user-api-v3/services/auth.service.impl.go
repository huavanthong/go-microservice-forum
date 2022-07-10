package services

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/huavanthong/microservice-golang/user-api-v3/models"
	"github.com/huavanthong/microservice-golang/user-api-v3/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AuthServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewAuthService(collection *mongo.Collection, ctx context.Context) AuthService {
	return &AuthServiceImpl{collection, ctx}
}

func (uc *AuthServiceImpl) SignUpUser(userInfo *models.SignUpInput) (*models.DBResponse, error) {

	/*** added the new user to the database with the InsertOne() function ***/
	var user models.User
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt
	user.Email = strings.ToLower(userInfo.Email)
	user.PasswordConfirm = ""
	user.Verified = true
	user.Role = "user"
	user.Photo = "default.png"
	user.Provider = "local"

	// security: hash password using bcrypt
	hashedPassword, _ := utils.HashPassword(userInfo.Password)
	user.Password = hashedPassword
	res, err := uc.collection.InsertOne(uc.ctx, &user)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("user with that email already exist")
		}
		return nil, err
	}

	/*** Create a unique index for the email field to ensure
	that no two users can have the same email address. ***/

	opt := options.Index()
	opt.SetUnique(true)

	// IndexModel represents a new index to be created.
	index := mongo.IndexModel{Keys: bson.M{"email": 1}, Options: opt}

	if _, err := uc.collection.Indexes().CreateOne(uc.ctx, index); err != nil {
		return nil, errors.New("could not create index for email")
	}

	var newUser *models.DBResponse
	query := bson.M{"_id": res.InsertedID}

	/*** find and return the user that was added to the database. ***/
	err = uc.collection.FindOne(uc.ctx, query).Decode(&newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (uc *AuthServiceImpl) SignInUser(*models.SignInInput) (*models.DBResponse, error) {
	return nil, nil
}
