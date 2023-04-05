package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/entities"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/interfaces/persistence"
)

// Define struct for MongoDB Basket Repository
type MongoDBBasketPersistence struct {
	client     *mongo.Client
	database   string
	collection string
}

// NewMongoDBBasketPersistence returns a new instance of MongoDBBasketPersistence
func NewMongoDBBasketPersistence(client *mongo.Client, database string, collection string) persistence.BasketPersistence {

	return &MongoDBBasketPersistence{
		client:     client,
		database:   database,
		collection: collection,
	}
}

func (mbp *MongoDBBasketPersistence) Create(userId string) (*entities.Basket, error) {

	// Create entity basket to storing in MongoDB
	var basket entities.Basket

	// Bson generate object id
	basket.ID = primitive.NewObjectID()
	basket.UserID = userId
	basket.Items = make([]entities.BasketItem, 0)
	basket.CreatedAt = time.Now()
	basket.UpdatedAt = basket.CreatedAt

	// Retrieves the MongoDB collection where the basket data is stored
	coll := mbp.client.Database(mbp.database).Collection(mbp.collection)

	// Insert to monodb
	_, err := coll.InsertOne(context.Background(), basket)
	if err != nil {
		return nil, fmt.Errorf("failed to create basket for user %s: %v", userId, err)
	}

	return &basket, nil
}

func (mbp *MongoDBBasketPersistence) GetByUserName(userName string) (*entities.Basket, error) {

	// Retrieves the MongoDB collection where the basket data is stored
	coll := mbp.client.Database(mbp.database).Collection(mbp.collection)

	// Create a filter to find the basket with the given user_name.
	filter := bson.M{"user_name": userName}

	// Search for the first document matching the filtembp.
	result := coll.FindOne(context.Background(), filter)
	if err := result.Err(); err != nil {
		return nil, fmt.Errorf("failed to get basket for user %s: %v", userName, err)
	}
	// If the search is successful, it decodes the document found into a entity
	basket := &entities.Basket{}
	if err := result.Decode(basket); err != nil {
		return nil, fmt.Errorf("failed to decode basket for user %s: %v", userName, err)
	}
	return basket, nil
}

func (mbp *MongoDBBasketPersistence) Update(basket *entities.Basket) (*entities.Basket, error) {
	// Retrieves the MongoDB collection where the basket data is stored
	coll := mbp.client.Database(mbp.database).Collection(mbp.collection)

	// Create a filter to find the basket with the given user_name.
	filter := bson.M{"user_name": basket.UserName}

	// Create an update statement with the $set operator, which sets
	// the items field to the new items in the Basket object.
	update := bson.M{"$set": bson.M{"items": basket.Items}}

	// Specify that the updated document should be returned
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	// Execute the update operation and retrieve the updated document.
	result := coll.FindOneAndUpdate(context.Background(), filter, update, opts)
	if err := result.Err(); err != nil {
		return nil, fmt.Errorf("failed to update basket for user %s: %v", basket.UserName, err)
	}
	// If success, decode from MongoDB response to basket entity
	if err := result.Decode(basket); err != nil {
		return nil, fmt.Errorf("failed to decode basket for user %s: %v", basket.UserName, err)
	}

	return basket, nil
}

func (mbp *MongoDBBasketPersistence) Delete(userName string) error {
	// Retrieves the MongoDB collection where the basket data is stored
	coll := mbp.client.Database(mbp.database).Collection(mbp.collection)

	// Specifies the shopping cart to be deleted based on the userName parametembp.
	filter := bson.M{"user_name": userName}

	// called on the collection to delete the shopping cart that matches the filtembp.
	_, err := coll.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("failed to delete basket for user %s: %v", userName, err)
	}
	return nil
}
