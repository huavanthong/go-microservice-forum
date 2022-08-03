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

type ProductServiceImpl struct {
	log        *zap.Logger
	collection *mongo.Collection
	ctx        context.Context
}

func NewProductServiceImpl(log *zap.Logger, collection *mongo.Collection, ctx context.Context) ProductService {
	return &ProductServiceImpl{log, collection, ctx}
}

func (p *ProductServiceImpl) CreateProduct(pr *payload.RequestCreateProduct) (*models.Product, error) {

	var temp models.Product
	temp.Name = pr.Name
	temp.Category = pr.Category
	temp.Summary = pr.Summary
	temp.Description = pr.Description
	temp.ImageFile = pr.ImageFile
	temp.Price = pr.Price
	temp.SKU = "test"
	temp.CreatedAt = time.Now()
	temp.UpdatedAt = temp.CreatedAt

	res, err := p.collection.InsertOne(p.ctx, pr)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("producvt with that title already exists")
		}
		return nil, err
	}

	opt := options.Index()
	opt.SetUnique(true)

	index := mongo.IndexModel{Keys: bson.M{"category": 1}, Options: opt}

	if _, err := p.collection.Indexes().CreateOne(p.ctx, index); err != nil {
		return nil, errors.New("could not create index for category")
	}

	var product *models.Product
	query := bson.M{"_id": res.InsertedID}
	if err = p.collection.FindOne(p.ctx, query).Decode(&product); err != nil {
		return nil, err
	}

	return product, nil

}

func (p *ProductServiceImpl) FindAllProducts(page int, limit int, currency string) ([]*models.Product, error) {

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

	// find all posts with optional data
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

	// if currency is empty, it return productList with the default of
	// base currency
	if currency == "" {
		return products, nil
	}

	// calculate exchange rate between base: Euro and dest: currency
	// rate, err := p.getRate(currency)
	// if err != nil {
	// 	p.log.Error("Unable to get rate", "currency", currency, "error", err)
	// }

	// // create a array to contain the rate products
	// pr := models.Product{}
	// // loop in productList to update to the product with rate
	// for _, p := range products {
	// 	// get a product
	// 	np := *p
	// 	// update it's currency with rate
	// 	np.Price = np.Price * rate
	// 	// push to a temp storage of product
	// 	pr = append(pr, &np)
	// }

	return products, nil
}

func (p *ProductServiceImpl) FindProductByID(id string, currency string) (*models.Product, error) {
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

	// if currency is empty, it return productList with the default of
	// base currency
	if currency == "" {
		return product, nil
	}

	return product, nil
}
func (p *ProductServiceImpl) FindProductByName(name string, currency string) (*models.Product, error) {

	// create container for data
	var product *models.Product

	// create a query command
	query := bson.M{"name": strings.ToLower(name)}

	// find one user by query command
	err := p.collection.FindOne(p.ctx, query).Decode(&product)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &models.Product{}, err
		}
		return nil, err
	}

	// if currency is empty, it return productList with the default of
	// base currency
	if currency == "" {
		return product, nil
	}

	return product, nil
}
func (p *ProductServiceImpl) FindProductByCategory(category string, currency string) (*models.Product, error) {

	// create container for data
	var product *models.Product

	// create a query command
	query := bson.M{"category": strings.ToLower(category)}

	// find one user by query command
	err := p.collection.FindOne(p.ctx, query).Decode(&product)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &models.Product{}, err
		}
		return nil, err
	}

	// if currency is empty, it return productList with the default of
	// base currency
	if currency == "" {
		return product, nil
	}

	return product, nil
}

func (p *ProductServiceImpl) UpdateProduct(id string, pr *models.Product) (*models.Product, error) {

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
func (p *ProductServiceImpl) DeleteProduct(id string) error {

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
