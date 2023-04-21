package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBClient struct {
	client   *mongo.Client
	database string
}

func NewMongoDBClient(connectionString string, database string) (*MongoDBClient, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return &MongoDBClient{
		client:   client,
		database: database,
	}, nil
}

func (mc *MongoDBClient) Disconnect() error {
	// Disconnect from MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := mc.client.Disconnect(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (mc *MongoDBClient) GetDatabase() *mongo.Database {
	// Get the database
	return mc.client.Database(mc.database)
}
