package migrations

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

func dropCollections(db *mongo.Database, ctx context.Context) error {

	// err := db.Collection("products").Drop()

	// err = db.Collection("category").Drop()

	// err = db.Collection("subcategory").Drop()

	return nil
}

func initCollections(db *mongo.Database, ctx context.Context) error {
	var err error
	// Specify the collections to migrate
	collections := map[string]MigrationHandler{
		"product": migrateProduct,
		//"category": migrateCategory,
		// "inventory": migrateInventory,
	}

	// Migrate each collection
	for name, handler := range collections {
		collection := db.Collection(name)
		err = handler(ctx, collection)
		if err != nil {
			log.Fatalf("Failed to migrate collection %s: %v", name, err)
		}
		fmt.Printf("Migrated collection %s\n", name)
	}

	return err
}

func HandleFlags(db *mongo.Database, ctx context.Context) error {

	cmdPtr := flag.String("cmd", "", "Command to execute (init-data, migrate)")
	colPtr := flag.String("col", "", "Collection to migrate")
	flag.Parse()

	switch *cmdPtr {
	case "init":
		// execute init-data command
		fmt.Println("Executing init-data command...")
		msg := ""
		if err := initCollections(db, ctx); err != nil {
			msg = fmt.Sprintf("Error initializing data in DB: %v", err)
		} else {
			msg = "Initialized data in DB."
		}

		fmt.Println(msg)
		return nil
	case "drop":
		// execute init-data command
		fmt.Println("Executing drop-data command...")

		msg := ""
		if err := dropCollections(db, ctx); err != nil {
			msg = fmt.Sprintf("Error dropping table: %v", err)
		} else {
			msg = "Dropped all collections in DB"
		}

		fmt.Println(msg)
		return nil

	case "migrate":
		// execute migrate command with specified collection
		if *colPtr == "" {

			return fmt.Errorf("Error: collection name is required for migrate command")
		}
		fmt.Printf("Executing migrate command for collection %s...\n", *colPtr)
		migrations(db, ctx, *colPtr)
		return nil

	default:
		return fmt.Errorf("Error: unknown command")
	}
}

// MigrationHandler is a function that migrates data for a specific collection.
type MigrationHandler func(ctx context.Context, collection *mongo.Collection) error

func migrations(db *mongo.Database, ctx context.Context, coll string) {

	// Specify the collections to migrate
	collections := map[string]MigrationHandler{
		"product":  migrateProduct,
		"category": migrateCategory,
		// "inventory": migrateInventory,
	}

	// Migrate each collection
	for name, handler := range collections {
		if name == coll {
			collection := db.Collection(name)
			err := handler(ctx, collection)
			if err != nil {
				log.Fatalf("Failed to migrate collection %s: %v", name, err)
			}
		}
		fmt.Printf("Migrated collection %s\n", name)
	}
}

// migrateProduct migrates data for the product collection.
func migrateProduct(ctx context.Context, collection *mongo.Collection) error {

	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return err
	}
	fmt.Println("Current working directory:", wd)

	// Read the data from the JSON file
	data, err := ioutil.ReadFile("product.json")
	if err != nil {
		return fmt.Errorf("Failed to read product data: %v", err)
	}

	// Unmarshal the JSON data into a slice of Product objects
	var products []interface{}
	err = json.Unmarshal(data, &products)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal product data: %v", err)
	}

	// Insert products into MongoDB collection
	productResult, err := collection.InsertMany(ctx, products)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal product data: %v", err)
	}

	fmt.Println("Products migrated successfully! %v", productResult)

	return nil
}

// migrateCategory migrates data for the category collection.
func migrateCategory(ctx context.Context, collection *mongo.Collection) error {
	// Read the data from the JSON file
	data, err := ioutil.ReadFile("category.json")
	if err != nil {
		return fmt.Errorf("Failed to read category data: %v", err)
	}

	// Unmarshal the JSON data into a slice of Category objects
	var categories []interface{}
	err = json.Unmarshal(data, &categories)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal category data: %v", err)
	}

	// Insert categories into the database
	categoryResult, err := collection.InsertMany(ctx, categories)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal product data: %v", err)
	}

	fmt.Println("Category migrated successfully! %v", categoryResult)
	return nil
}

// migrateCategory migrates data for the category collection.
func migrateInventory(ctx context.Context, collection *mongo.Collection) error {
	// Read the data from the JSON file
	data, err := ioutil.ReadFile("category.json")
	if err != nil {
		return fmt.Errorf("Failed to read category data: %v", err)
	}

	// Unmarshal the JSON data into a slice of Category objects
	var categories []interface{}
	err = json.Unmarshal(data, &categories)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal category data: %v", err)
	}

	// Insert categories into the database
	categoryResult, err := collection.InsertMany(ctx, categories)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal product data: %v", err)
	}

	fmt.Println("Category migrated successfully! %v", categoryResult)
	return nil
}
