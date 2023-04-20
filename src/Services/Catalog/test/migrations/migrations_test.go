package migrations

import (
	"context"
	"fmt"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/migrations"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db   *mongo.Database
	ctx  context.Context
	host = "localhost"
	port = "27017"
)

func TestMain(m *testing.M) {
	// Set up a MongoDB memory cache for testing.
	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", host, port)))
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	db = client.Database("test")

	// Run tests
	exitVal := m.Run()

	// Clean up
	if err := db.Drop(ctx); err != nil {
		log.Fatalf("Failed to drop test database: %v", err)
	}

	os.Exit(exitVal)
}

func TestDropCollections(t *testing.T) {

	// Create fixture
	collections := []string{"collection1", "collection2", "collection3"}
	for _, coll := range collections {
		db.Collection(coll).InsertOne(context.Background(), bson.M{})
	}

	// Set the command line arguments
	os.Args = []string{"cmd", "-cmd", "drop"}

	// Run test
	err := migrations.HandleFlags(db, ctx)
	assert.NoError(t, err)

	for _, coll := range collections {
		count, _ := db.Collection(coll).CountDocuments(context.Background(), bson.M{})
		assert.Equal(t, int64(0), count)
	}

}

func TestInitCollections(t *testing.T) {
	// Set the command line arguments
	os.Args = []string{"cmd", "-cmd", "init"}

	// Run test
	err := migrations.HandleFlags(db, ctx)
	assert.NoError(t, err)
}

func TestMigrationCollections(t *testing.T) {
	// Set the command line arguments
	os.Args = []string{"cmd", "-cmd", "drop", "-coll", "product"}

	// Run test
	err := migrations.HandleFlags(db, ctx)
	assert.NoError(t, err)
}
