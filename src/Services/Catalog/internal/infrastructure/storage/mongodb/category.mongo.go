package mongodb

import (
	"context"
	"errors"
	"strings"
	"time"

	"go.uber.org/zap"

	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/domain/entities"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CategoryStorage struct {
	log        *zap.Logger
	collection *mongo.Collection
	ctx        context.Context
}

func NewCategoryStorage(log *zap.Logger, collection *mongo.Collection, ctx context.Context) *CategoryStorage {
	return &CategoryStorage{
		log,
		collection,
		ctx,
	}
}

func (cr *CategoryStorage) Create(c *entities.Category) (*entities.Category, error) {

	cr.log.Info("Creating a new category...")

	// Initialize the basic info of category
	c.CreatedAt = time.Now().Format(time.RFC3339)
	c.UpdatedAt = c.CreatedAt

	cr.log.Info("Generating ObjectID...")
	/*** ObjectID: Bson generate object id ***/
	c.ID = primitive.NewObjectID()

	cr.log.Info("Inserting category into database...")
	_, err := cr.collection.InsertOne(cr.ctx, c)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			cr.log.Error("Category with that ccode already exists", zap.Error(err))
			return nil, errors.New("category with that ccode already exists")
		}
		cr.log.Error("Error inserting category into database", zap.Error(err))
		return nil, err
	}

	cr.log.Info("Creating index for ccode...")
	// Create Indexesfor pcode, it help you easy to find product by pcode
	opt := options.Index()
	opt.SetUnique(true)

	index := mongo.IndexModel{Keys: bson.M{"ccode": 1}, Options: opt}

	if _, err := cr.collection.Indexes().CreateOne(cr.ctx, index); err != nil {
		cr.log.Error("Error creating index for ccode", zap.Error(err))
		return nil, errors.New("could not create index for ccode")
	}

	var category *entities.Category
	// query := bson.M{"_id": res.InsertedID}
	query := bson.M{"_id": c.ID}

	cr.log.Info("Fetching category...")
	if err = cr.collection.FindOne(cr.ctx, query).Decode(&category); err != nil {
		cr.log.Error("Error fetching category", zap.Error(err))
		return nil, err
	}

	cr.log.Info("Category created successfully")
	return category, nil

}

func (cr *CategoryStorage) GetByID(id string) (*entities.Category, error) {

	// convert string id to objectID
	obId, _ := primitive.ObjectIDFromHex(id)

	// create a query command by id
	query := bson.M{"_id": obId}

	// create container
	var category *entities.Category

	// find one post by query command
	if err := cr.collection.FindOne(cr.ctx, query).Decode(&category); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no document with that Id exists")
		}

		return nil, err
	}

	return category, nil
}

func (cr *CategoryStorage) Update(c *entities.Category) (*entities.Category, error) {

	doc, err := utils.ToDoc(c)
	if err != nil {
		return nil, err
	}

	query := bson.D{{Key: "_id", Value: c.ID}}
	update := bson.D{{Key: "$set", Value: doc}}
	res := cr.collection.FindOneAndUpdate(cr.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatedCategory *entities.Category

	if err := res.Decode(&updatedCategory); err != nil {
		return nil, errors.New("no post with that Id exists")
	}

	return updatedCategory, nil
}

func (cr *CategoryStorage) Delete(id string) error {

	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	res, err := cr.collection.DeleteOne(cr.ctx, query)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("no document with that Id exists")
	}

	return nil
}

func (cr *CategoryStorage) GetCategories(filter *entities.CategoryFilter, pagination *entities.Pagination) ([]*entities.Category, error) {

	// create a query command
	query := bson.M{}

	// Check condition
	opts := options.Find()
	if pagination != nil {
		opts.SetSkip(pagination.Offset())
		opts.SetLimit(pagination.Limit)
	}

	// find all categories with optional data
	cursor, err := cr.collection.Find(cr.ctx, query, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(cr.ctx)

	// create container for data
	var categories []*entities.Category

	// with data find out, we will decode them and append to array
	for cursor.Next(cr.ctx) {
		category := &entities.Category{}
		err := cursor.Decode(category)

		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	// if any item error, return err
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// if data is empty, return nil
	if len(categories) == 0 {
		return []*entities.Category{}, nil
	}

	return categories, nil
}

func (c *CategoryStorage) FindCategoryByName(name string) ([]*entities.Category, error) {

	// we should create query option

	// create a query command
	query := bson.M{"name": strings.ToLower(name)}

	// find one user by query command
	cursor, err := c.collection.Find(c.ctx, query, nil)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(c.ctx)

	// create container for data
	var categories []*entities.Category

	// with data find out, we will decode them and append to array
	for cursor.Next(c.ctx) {
		category := &entities.Category{}
		err := cursor.Decode(category)

		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	// if any item error, return err
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// if data is empty, return nil
	if len(categories) == 0 {
		return []*entities.Category{}, nil
	}

	return categories, nil
}
