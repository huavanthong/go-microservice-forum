package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/entities"
)

type MongoDBBasketRepository struct {
	client     *mongo.Client
	database   string
	collection string
}

func NewMongoDBBasketRepository(client *mongo.Client, database string, collection string) (*MongoDBBasketRepository, error) {

	return &MongoDBBasketRepository{
		client:     client,
		database:   database,
		collection: collection,
	}, nil
}

func (r *MongoDBBasketRepository) Session() *mongo.Session {
	return r.client.StartSession()
}

func (r *MongoDBBasketRepository) GetBasket(userName string) (*entities.ShoppingCart, error) {

	// Get collection
	coll := r.client.Database(r.database).Collection(r.collection)

	filter := map[string]interface{}{
		"user_name": userName,
	}

	result := coll.FindOne(context.Background(), filter)
	if err := result.Err(); err != nil {
		return nil, fmt.Errorf("failed to get basket for user %s: %v", userName, err)
	}

	basket := &entities.ShoppingCart{}
	if err := result.Decode(basket); err != nil {
		return nil, fmt.Errorf("failed to decode basket for user %s: %v", userName, err)
	}

	return basket, nil
}

func (r *MongoDBBasketRepository) UpdateBasket(basket *entities.ShoppingCart) (*entities.ShoppingCart, error) {
	session := r.Session()
	defer session.Close()

	c := session.DB(r.database).C(r.collection)

	err := c.Update(bson.M{"username": basket.Username}, basket)
	if err != nil {
		return nil, fmt.Errorf("could not update basket for user %s: %v", basket.Username, err)
	}

	return basket, nil
}

func (r *MongoDBBasketRepository) DeleteBasket(userName string) error {
	session := r.Session()
	defer session.Close()

	c := session.DB(r.database).C(r.collection)

	err := c.Remove(bson.M{"username": userName})
	if err != nil {
		return fmt.Errorf("could not delete basket for user %s: %v", userName, err)
	}

	return nil
}
