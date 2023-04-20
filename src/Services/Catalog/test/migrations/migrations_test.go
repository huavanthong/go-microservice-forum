package migrations

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
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
	host = "catalogdb"
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
	err := dropCollections(db, ctx)
	assert.NoError(t, err)
}

func TestInitCollections(t *testing.T) {
	// Initialize data for Category
	jsonFile, err := os.Open("category.json")
	assert.NoError(t, err)
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var categories []interface{}
	if err := json.Unmarshal(byteValue, &categories); err != nil {
		t.Fatalf("Failed to unmarshal category data: %v", err)
	}

	// Insert category data directly into the database
	categoryResult, err := db.Collection("category").InsertMany(ctx, categories)
	if err != nil {
		t.Fatalf("Failed to insert category data: %v", err)
	}

	// Run test
	err = initCollections(db, ctx)
	assert.NoError(t, err)

	// Verify that the number of inserted documents is the same as the number of documents inserted directly into the database
	c, err := db.Collection("category").CountDocuments(ctx, bson.M{})
	if err != nil {
		t.Fatalf("Failed to count documents: %v", err)
	}

	assert.Equal(t, int64(len(categories)), c)
	assert.Equal(t, categoryResult.InsertedIDs, c)
}

func TestHandleFlag(t *testing.T) {
	// Test enable dummy data
	os.Args = []string{"", "-enable-data", "-init-data"}
	handleFlag(db, ctx)

	c, err
