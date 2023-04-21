package mongodb

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/entities"

	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/interfaces/persistence"
)

// Define struct for MongoDB Basket Repository
type MongoDBBasketPersistence struct {
	logger     *logrus.Entry // add logger
	collection *mongo.Collection
	ctx        context.Context
}

// NewMongoDBBasketPersistence returns a new instance of MongoDBBasketPersistence
func NewMongoDBBasketPersistence(logger *logrus.Entry, collection *mongo.Collection, ctx context.Context) persistence.BasketPersistence {

	return &MongoDBBasketPersistence{
		logger:     logrus.WithField("module", "MongoDBBasketPersistence"),
		collection: collection,
		ctx:        ctx,
	}
}

func (mbp *MongoDBBasketPersistence) Create(basket *entities.Basket) (*entities.Basket, error) {

	// Log the user ID being processed
	mbp.logger.Info("Creating basket by user id %s in MongoDB ", basket.UserID)

	// Insert to monodb
	_, err := mbp.collection.InsertOne(context.Background(), basket)
	if err != nil {
		return nil, fmt.Errorf("failed to create basket for user %s: %v", basket.UserID, err)
	}

	mbp.logger.Info("Created basket %+v for user ID %s in MongoDB", basket, basket.UserID)

	return basket, nil
}

func (mbp *MongoDBBasketPersistence) Get(userId string) (*entities.Basket, error) {

	// Log the user ID being processed
	mbp.logger.Info("Getting basket by user id %s in MongoDB ", userId)

	// Create a filter to find the basket with the given user_name.
	filter := bson.M{"user_id": userId}

	// Search for the first document matching the filtembp.
	result := mbp.collection.FindOne(context.Background(), filter)
	if err := result.Err(); err != nil {
		return nil, fmt.Errorf("failed to get basket for user %s: %v", userId, err)
	}
	// If the search is successful, it decodes the document found into a entity
	basket := &entities.Basket{}
	if err := result.Decode(basket); err != nil {
		return nil, fmt.Errorf("failed to decode basket for user %s: %v", userId, err)
	}

	mbp.logger.Info("Got basket %+v for user ID %s", basket, userId)

	return basket, nil
}

func (mbp *MongoDBBasketPersistence) Update(basket *entities.Basket) (*entities.Basket, error) {

	// Log the user ID being processed
	mbp.logger.Info("Updating basket with info: ", basket)

	// Create a filter to find the basket with the given user_name.
	filter := bson.M{"user_id": basket.UserID}

	// Create an update statement with the $set operator, which sets
	// the items field to the new items in the Basket object.
	update := bson.M{"$set": bson.M{"items": basket.Items}}

	// Specify that the updated document should be returned
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	// Execute the update operation and retrieve the updated document.
	result := mbp.collection.FindOneAndUpdate(context.Background(), filter, update, opts)
	if err := result.Err(); err != nil {
		return nil, fmt.Errorf("failed to update basket for user %s: %v", basket.UserID, err)
	}

	var updatedBasket *entities.Basket
	// If success, decode from MongoDB response to basket entity
	if err := result.Decode(&updatedBasket); err != nil {
		return nil, fmt.Errorf("failed to decode basket for user %s: %v", basket.UserID, err)
	}

	// Log the basket being returned
	mbp.logger.Info("Updated basket %+v for user ID %s", updatedBasket, updatedBasket.UserID)

	return basket, nil
}

func (mbp *MongoDBBasketPersistence) Delete(userId string) error {

	// Log the user ID being processed
	mbp.logger.Info("Deleting basket by user id %s in MongoDB ", userId)

	// Specifies the shopping cart to be deleted based on the userName parametembp.
	filter := bson.M{"user_id": userId}

	// called on the collection to delete the shopping cart that matches the filtembp.
	_, err := mbp.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("failed to delete basket for user %s: %v", userId, err)
	}

	mbp.logger.Info("Deleted basket for user id %s in MongoDB", userId)

	return nil
}
