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

	// Read data from JSON file
	jsonFile, err := os.Open("products.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened products.json")
	defer jsonFile.Close()

	// Initialize data for Category
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var products []interface{}
	if err := json.Unmarshal(byteValue, &products); err != nil {
		return err
	}

	productResult, err := db.Collection("category").InsertMany(ctx, products)
	fmt.Printf("Inserted %v documents into product collection!\n", productResult)

	return err
}

func handleFlag(db *mongo.Database, ctx context.Context) {

	enableDummyData := flag.Bool("enable-data", false, "Enable dummy data for testing")
	initData := flag.Bool("init-data", false, "Set this flag if DB should be initialized with dummy data")
	drop := flag.Bool("drop-table", false, "Set this flag if you wan't to drop all user data in your DB")
	flag.Parse()

	if *enableDummyData {
		if *drop {
			msg := ""
			if err := dropCollections(db, ctx); err != nil {
				msg = fmt.Sprintf("Error dropping table: %v", err)
			} else {
				msg = "Dropped all collections in DB"
			}

			fmt.Println(msg)
			os.Exit(0)
		}

		if *initData {
			msg := ""
			if err := initCollections(db, ctx); err != nil {
				msg = fmt.Sprintf("Error initializing data in DB: %v", err)
			} else {
				msg = "Initialized data in DB."
			}

			fmt.Println(msg)
			os.Exit(0)
		}
	}

}

// MigrationHandler is a function that migrates data for a specific collection.
type MigrationHandler func(ctx context.Context, collection *mongo.Collection) error

func Migrations(db *mongo.Database, ctx context.Context) {

	// Specify the collections to migrate
	collections := map[string]MigrationHandler{
		"product":  migrateProduct,
		"category": migrateCategory,
		// "inventory": migrateInventory,
	}

	// Migrate each collection
	for name, handler := range collections {
		collection := db.Collection(name)
		err := handler(ctx, collection)
		if err != nil {
			log.Fatalf("Failed to migrate collection %s: %v", name, err)
		}
		fmt.Printf("Migrated collection %s\n", name)
	}
}

// migrateProduct migrates data for the product collection.
func migrateProduct(ctx context.Context, collection *mongo.Collection) error {
	// Read the data from the JSON file
	data, err := ioutil.ReadFile("products.json")
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
}
