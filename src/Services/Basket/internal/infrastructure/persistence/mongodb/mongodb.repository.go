package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/entities"
)

// Define struct for MongoDB Basket Repository
type MongoDBBasketRepository struct {
	client     *mongo.Client
	database   string
	collection string
}

// NewMongoDBBasketRepository returns a new instance of MongoDBBasketRepository
func NewMongoDBBasketRepository(client *mongo.Client, database string, collection string) *MongoDBBasketRepository {

	return &MongoDBBasketRepository{
		client:     client,
		database:   database,
		collection: collection,
	}
}

func (r *MongoDBBasketRepository) Create(userName string) (*entities.ShoppingCart, error) {

	// Create entity for Shopping Cart start for shopping
	cart := &entities.ShoppingCart{
		UserName: userName,
		Items:    make([]entities.ShoppingCartItem, 0),
	}

	// Retrieves the MongoDB collection where the basket data is stored
	coll := r.client.Database(r.database).Collection(r.collection)

	// Insert to monodb
	_, err := coll.InsertOne(context.Background(), cart)
	if err != nil {
		return nil, fmt.Errorf("failed to create basket for user %s: %v", userName, err)
	}

	return cart, nil
}

func (r *MongoDBBasketRepository) GetByUserName(userName string) (*entities.ShoppingCart, error) {

	// Retrieves the MongoDB collection where the basket data is stored
	coll := r.client.Database(r.database).Collection(r.collection)

	// Create a filter to find the basket with the given user_name.
	filter := bson.M{"user_name": userName}

	// Search for the first document matching the filter.
	result := coll.FindOne(context.Background(), filter)
	if err := result.Err(); err != nil {
		return nil, fmt.Errorf("failed to get basket for user %s: %v", userName, err)
	}
	// If the search is successful, it decodes the document found into a entity
	basket := &entities.ShoppingCart{}
	if err := result.Decode(basket); err != nil {
		return nil, fmt.Errorf("failed to decode basket for user %s: %v", userName, err)
	}
	return basket, nil
}

func (r *MongoDBBasketRepository) Update(basket *entities.ShoppingCart) (*entities.ShoppingCart, error) {
	// Retrieves the MongoDB collection where the basket data is stored
	coll := r.client.Database(r.database).Collection(r.collection)

	// Create a filter to find the basket with the given user_name.
	filter := bson.M{"user_name": basket.UserName}

	// Create an update statement with the $set operator, which sets
	// the items field to the new items in the ShoppingCart object.
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

func (r *MongoDBBasketRepository) Delete(userName string) error {
	// Retrieves the MongoDB collection where the basket data is stored
	coll := r.client.Database(r.database).Collection(r.collection)

	// Specifies the shopping cart to be deleted based on the userName parameter.
	filter := bson.M{"user_name": userName}

	// called on the collection to delete the shopping cart that matches the filter.
	_, err := coll.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("failed to delete basket for user %s: %v", userName, err)
	}
	return nil
}
