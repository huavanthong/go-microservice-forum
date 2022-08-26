package services

import (
	"context"
	"errors"
	"strings"
	"time"

	"go.uber.org/zap"

	"github.com/huavanthong/microservice-golang/product-api-v3/models"
	"github.com/huavanthong/microservice-golang/product-api-v3/payload"
	"github.com/huavanthong/microservice-golang/product-api-v3/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CategoryServiceImpl struct {
	log        *zap.Logger
	collection *mongo.Collection
	ctx        context.Context
}

func NewCategoryServiceImpl(log *zap.Logger, collection *mongo.Collection, ctx context.Context) CategoryService {
	return &CategoryServiceImpl{log, collection, ctx}
}

func (c *CategoryServiceImpl) CreateCategory(pr *payload.RequestCreateCategory) (*models.Category, error) {

	// Initialize the basic info of product
	var temp models.Product
	temp.Name = pr.Name
	temp.ProductType = pr.Name
	temp.Category = pr.Category
	temp.Summary = pr.Summary
	temp.Description = pr.Description
	temp.ImageFile = pr.ImageFile
	temp.Price = pr.Price
	temp.ProductCode = "p" + utils.RandCode(9)
	temp.SKU = "ABC-XXX-YYY"
	temp.CreatedAt = time.Now().String()
	temp.UpdatedAt = temp.CreatedAt

	/*** ObjectID: Bson generate object id ***/
	temp.ID = primitive.NewObjectID()

	_, err := p.collection.InsertOne(p.ctx, temp)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("product with that pcode already exists")
		}
		return nil, err
	}
	// Create Indexesfor pcode, it help you easy to find product by pcode
	opt := options.Index()
	opt.SetUnique(true)

	index := mongo.IndexModel{Keys: bson.M{"pcode": 1}, Options: opt}

	if _, err := p.collection.Indexes().CreateOne(p.ctx, index); err != nil {
		return nil, errors.New("could not create index for pcode")
	}

	var product *models.Product
	// query := bson.M{"_id": res.InsertedID}
	query := bson.M{"_id": temp.ID}

	if err = p.collection.FindOne(p.ctx, query).Decode(&product); err != nil {
		return nil, err
	}

	return product, nil

}

func (c *CategoryServiceImpl) FindAllCategories(page int, limit int) ([]*models.Category, error) {

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

	// find all products with optional data
	cursor, err := p.collection.Find(p.ctx, query, &opt)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(p.ctx)

	// create container for data
	var products []*models.Product

	// with data find out, we will decode them and append to array
	for cursor.Next(p.ctx) {
		product := &models.Product{}
		err := cursor.Decode(product)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	// if any item error, return err
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// if data is empty, return nil
	if len(products) == 0 {
		return []*models.Product{}, nil
	}

	return products, nil
}

func (c *CategoryServiceImpl) FindCategoryByID(id string) (*models.Category, error) {
	// convert string id to objectID
	obId, _ := primitive.ObjectIDFromHex(id)

	// create a query command by id
	query := bson.M{"_id": obId}

	// create container
	var product *models.Product

	// find one post by query command
	if err := p.collection.FindOne(p.ctx, query).Decode(&product); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no document with that Id exists")
		}

		return nil, err
	}

	return product, nil
}

func (c *CategoryServiceImpl) FindCategoryByName(name string) ([]*models.Category, error) {

	// we should create query option

	// create a query command
	query := bson.M{"name": strings.ToLower(name)}

	// find one user by query command
	cursor, err := p.collection.Find(p.ctx, query, nil)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(p.ctx)

	// create container for data
	var products []*models.Product

	// with data find out, we will decode them and append to array
	for cursor.Next(p.ctx) {
		product := &models.Product{}
		err := cursor.Decode(product)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	// if any item error, return err
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// if data is empty, return nil
	if len(products) == 0 {
		return []*models.Product{}, nil
	}

	return products, nil
}

func (c *CategoryServiceImpl) UpdateCategory(id string, pr *payload.RequestUpdateProduct) (*models.Product, error) {

	doc, err := utils.ToDoc(pr)
	if err != nil {
		return nil, err
	}

	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.D{{Key: "_id", Value: obId}}
	update := bson.D{{Key: "$set", Value: doc}}
	res := p.collection.FindOneAndUpdate(p.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatedPost *models.Product

	if err := res.Decode(&updatedPost); err != nil {
		return nil, errors.New("no post with that Id exists")
	}

	return updatedPost, nil
}

func (c *CategoryServiceImpl) DeleteCategory(id string) error {

	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	res, err := p.collection.DeleteOne(p.ctx, query)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("no document with that Id exists")
	}

	return nil
}
