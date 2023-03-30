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

type ProductStorage struct {
	log        *zap.Logger
	collection *mongo.Collection
	ctx        context.Context
}

func NewProductStorage(log *zap.Logger, collection *mongo.Collection, ctx context.Context) *ProductStorage {
	return &ProductStorage{
		log,
		collection,
		ctx,
	}
}

func (ps *ProductStorage) Create(p *entities.Product) (*entities.Product, error) {

	// Create date
	p.CreatedAt = time.Now().Format(time.RFC3339)
	p.UpdatedAt = p.CreatedAt

	// Bson generate object id
	p.ID = primitive.NewObjectID()

	// Insert product to mongodb
	_, err := p.collection.InsertOne(ps.ctx, p)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("product with that pcode already exists")
		}
		return nil, err
	}

	// Create Indexes for pcode, it help you easy to find product by pcode
	opt := options.Index()
	opt.SetUnique(true)

	index := mongo.IndexModel{Keys: bson.M{"pcode": 1}, Options: opt}

	if _, err := p.collection.Indexes().CreateOne(p.ctx, index); err != nil {
		return nil, errors.New("could not create index for pcode")
	}

	var product *entities.Product
	// query := bson.M{"_id": res.InsertedID}
	query := bson.M{"_id": p.ID}

	if err = p.collection.FindOne(p.ctx, query).Decode(&product); err != nil {
		return nil, err
	}

	return product, nil
}

func (ps *ProductStorage) GetByID(id string) (*entities.Product, error) {

	// convert string id to objectID
	obId, _ := primitive.ObjectIDFromHex(id)

	// create a query command by id
	query := bson.M{"_id": obId}

	// create container
	var product *entities.Product

	// find one post by query command
	if err := ps.collection.FindOne(ps.ctx, query).Decode(&product); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no document with that Id exists")
		}

		return nil, err
	}

	return product, nil
}

func (ps *ProductStorage) Update(p *entities.Product) (*entities.Product, error) {

	doc, err := utils.ToDoc(p)
	if err != nil {
		return nil, err
	}

	// Get object ID
	obId, _ := primitive.ObjectIDFromHex(p.ID)

	query := bson.D{{Key: "_id", Value: obId}}
	update := bson.D{{Key: "$set", Value: doc}}
	res := p.collection.FindOneAndUpdate(p.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatedPost *entities.Product

	if err := res.Decode(&updatedPost); err != nil {
		return nil, errors.New("no post with that Id exists")
	}

	return updatedPost, nil
}

func (ps *ProductStorage) Delete(id string) error {

	// convert string id to objectID
	obId, _ := primitive.ObjectIDFromHex(id)

	query := bson.M{"_id": obId}

	res, err := ps.collection.DeleteOne(ps.ctx, query)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("no document with that Id exists")
	}

	return nil
}

func (ps *ProductStorage) GetProducts(filter *entities.ProductFilter, pagination *entities.Pagination) ([]*entities.Product, int64, error) {

	query := bson.M{}
	// Check filter condition
	if filter != nil {
		if filter.Category != "" {
			query["category.name"] = filter.Category
		}
		if filter.ProductType != "" {
			query["ptype"] = filter.ProductType
		}
		if filter.Brand != "" {
			query["brand.name"] = filter.Brand
		}
		if filter.PriceMin > 0 {
			query["price"] = bson.M{"$gte": filter.PriceMin}
		}
		if filter.PriceMax > 0 {
			if query["price"] != nil {
				query["price"].(bson.M)["$lte"] = filter.PriceMax
			} else {
				query["price"] = bson.M{"$lte": filter.PriceMax}
			}
		}
	}

	// Count product by query
	var count int64
	count, err := ps.collection.CountDocuments(ps.ctx, query)
	if err != nil {
		return nil, errors.New(err, "failed to count products in DB")
	}

	// Check condition
	opts := options.Find()
	if pagination != nil {
		opts.SetSkip(pagination.Offset())
		opts.SetLimit(pagination.Limit)
	}
	cursor, err := ps.collection.Find(ps.ctx, query, opts)
	if err != nil {
		return nil, errors.New(err, "failed to find products in DB")
	}
	defer cursor.Close(ps.ctx)

	// Create container for data
	var products []*entities.Product

	// Found data, then decode them and append to array
	for cursor.Next(ps.ctx) {

		var product entities.Product

		if err := cursor.Decode(&product); err != nil {
			return nil, errors.New(err, "failed to decode product")
		}

		products = append(products, &product)
	}

	// if any item error, return err
	if err := cursor.Err(); err != nil {
		return nil, errors.New(err, "failed to iterate products")
	}

	// Warning: if data is empty, return nil
	if len(products) == 0 {
		return []*entities.Product{}, nil
	}

	return products, nil
}

func (p *ProductStorage) FindProductByName(name string, currency string) ([]*entities.Product, error) {

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
	var products []*entities.Product

	// with data find out, we will decode them and append to array
	for cursor.Next(p.ctx) {
		product := &entities.Product{}
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
		return []*entities.Product{}, nil
	}

	// if currency is empty, it return productList with the default of
	// base currency
	if currency == "" {
		return products, nil
	}

	return products, nil
}

func (p *ProductStorage) FindProductByCategory(category string, currency string) ([]*entities.Product, error) {

	// we should create query option

	// create a query command
	// query := bson.M{"category": strings.ToLower(category)}
	// fmt.Println("Check 1: ", query)
	query := bson.D{{"category", category}}

	// find one user by query command
	cursor, err := p.collection.Find(p.ctx, query, nil)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(p.ctx)

	// // Find all documents in which the "name" field is "Bob".
	// // Specify the Sort option to sort the returned documents by age in
	// // ascending order.
	// opts := options.Find().SetSort(bson.D{{"age", 1}})
	// cursor, err := p.Find(context.TODO(), bson.D{{"name", "Bob"}}, opts)
	// if err != nil {
	// 	return nil, err
	// }

	// create container for data
	var products []*entities.Product

	// with data find out, we will decode them and append to array
	for cursor.Next(p.ctx) {
		product := &entities.Product{}
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
		return []*entities.Product{}, nil
	}

	// if currency is empty, it return productList with the default of
	// base currency
	if currency == "" {
		return products, nil
	}

	return products, nil
}
