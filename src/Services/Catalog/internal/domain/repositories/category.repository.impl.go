package repositories

import (
	"context"
	"errors"
	"strings"
	"time"

	"go.uber.org/zap"

	"github.com/huavanthong/microservice-golang/product-api-v3/utils"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/api/models"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/domain/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CategoryRepositoryImpl struct {
	log        *zap.Logger
	collection *mongo.Collection
	ctx        context.Context
}

func NewCategoryRepositoryImpl(log *zap.Logger, collection *mongo.Collection, ctx context.Context) CategoryRepository {
	return &CategoryRepositoryImpl{log, collection, ctx}
}

func (c *CategoryRepositoryImpl) CreateCategory(rc *models.RequestCreateCategory) (*entities.Category, error) {

	// Initialize the basic info of category
	var temp entities.Category
	temp.Name = rc.Name
	temp.CategoryCode = "c" + utils.RandCode(4)
	temp.SubCategories = nil
	temp.Description = rc.Description
	temp.CreatedAt = time.Now().String()
	temp.UpdatedAt = temp.CreatedAt

	/*** ObjectID: Bson generate object id ***/
	temp.ID = primitive.NewObjectID()

	_, err := c.collection.InsertOne(c.ctx, temp)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("category with that ccode already exists")
		}
		return nil, err
	}
	// Create Indexesfor pcode, it help you easy to find product by pcode
	opt := options.Index()
	opt.SetUnique(true)

	index := mongo.IndexModel{Keys: bson.M{"ccode": 1}, Options: opt}

	if _, err := c.collection.Indexes().CreateOne(c.ctx, index); err != nil {
		return nil, errors.New("could not create index for ccode")
	}

	var category *entities.Category
	// query := bson.M{"_id": res.InsertedID}
	query := bson.M{"_id": temp.ID}

	if err = c.collection.FindOne(c.ctx, query).Decode(&category); err != nil {
		return nil, err
	}

	return category, nil

}

func (c *CategoryRepositoryImpl) FindAllCategories(page int, limit int) ([]*entities.Category, error) {

	// page return product
	if page == 0 {
		page = 1
	}

	// limit data return
	if limit == 0 {
		limit = 20
	}

	skip := (page - 1) * limit

	opt := options.FindOptions{}
	opt.SetLimit(int64(limit))
	opt.SetSkip(int64(skip))

	// create a query command
	query := bson.M{}

	// find all categories with optional data
	cursor, err := c.collection.Find(c.ctx, query, &opt)
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

func (c *CategoryRepositoryImpl) FindCategoryByID(id string) (*entities.Category, error) {
	// convert string id to objectID
	obId, _ := primitive.ObjectIDFromHex(id)

	// create a query command by id
	query := bson.M{"_id": obId}

	// create container
	var category *entities.Category

	// find one post by query command
	if err := c.collection.FindOne(c.ctx, query).Decode(&category); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no document with that Id exists")
		}

		return nil, err
	}

	return category, nil
}

func (c *CategoryRepositoryImpl) FindCategoryByName(name string) ([]*entities.Category, error) {

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

func (c *CategoryRepositoryImpl) UpdateCategory(id string, pr *models.RequestUpdateCategory) (*entities.Category, error) {

	doc, err := utils.ToDoc(pr)
	if err != nil {
		return nil, err
	}

	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.D{{Key: "_id", Value: obId}}
	update := bson.D{{Key: "$set", Value: doc}}
	res := c.collection.FindOneAndUpdate(c.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatedCategory *entities.Category

	if err := res.Decode(&updatedCategory); err != nil {
		return nil, errors.New("no post with that Id exists")
	}

	return updatedCategory, nil
}

func (c *CategoryRepositoryImpl) DeleteCategory(id string) error {

	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	res, err := c.collection.DeleteOne(c.ctx, query)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("no document with that Id exists")
	}

	return nil
}
