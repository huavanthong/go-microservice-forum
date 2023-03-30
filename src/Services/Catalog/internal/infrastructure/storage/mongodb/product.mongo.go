package mongodb

import (
	"context"
	"errors"
	"strings"
	"time"

	"go.uber.org/zap"

	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/api/models"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/domain/entities"
	"github.com/huavanthong/microservice-golang/src/Services/Catalog/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductRepository struct {
	log        *zap.Logger
	collection *mongo.Collection
	ctx        context.Context
}

func NewProductRepository(log *zap.Logger, collection *mongo.Collection, ctx context.Context) *ProductRepository {
	return &ProductRepository{
		log,
		collection,
		ctx,
	}
}

func (p *ProductRepository) CreateProduct(pr *models.RequestCreateProduct) (*entities.Product, error) {

	// Use Factory Design Pattern to get product following product type
	// Implement later
	// temp, _ := entities.GetProductType(entities.ProductType(pr.ProductType))

	temp := entities.Product{}
	// Initialize the basic info of product
	temp.SetName(pr.Name)
	temp.SetCategory(pr.Category)
	temp.SetSummary(pr.Summary)
	temp.SetDescription(pr.Description)
	temp.SetImageFile(pr.ImageFile)
	temp.SetPrice(pr.Price)
	temp.SetProductCode("p" + utils.RandCode(9))
	temp.SetSKU("ABC-XXX-YYY")
	temp.SetCreatedAt(time.Now().String())
	temp.SetUpdatedAt(temp.GetCreatedAt())

	/*** ObjectID: Bson generate object id ***/
	temp.SetID(primitive.NewObjectID())

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

	var product *entities.Product
	// query := bson.M{"_id": res.InsertedID}
	query := bson.M{"_id": temp.GetID()}

	if err = p.collection.FindOne(p.ctx, query).Decode(&product); err != nil {
		return nil, err
	}

	return product, nil

}

// func (p *ProductRepository) CreateProductPhone(pr *models.RequestCreateProduct) (*entities.Product_phone, error) {

// 	// Use Factory Design Pattern to get product following product type
// 	productType, perr := entities.GetProductType(entities.ProductType(pr.ProductType))
// 	if perr != nil {
// 		return nil, perr
// 	}

// 	switch utils.TypeOfModel(productType) {
// 	case "phone":
// 		productPhone, _ := productType.(*entities.Product_phone)
// 		break
// 	case "dien-tu":
// 		productDienTu, _ := productType.(*entities.Product_dientu)
// 		break
// 	case "thoi-trang":
// 		productThoiTrang, _ := productType.(*entities.Product_thoitrang)
// 	default:
// 		return nil, fmt.Errorf("Wrong product type passed")
// 	}

// 	// Initialize the basic info of product
// 	productPhone.SetName(pr.Name)
// 	productPhone.SetCategory(pr.Category)
// 	productPhone.SetSummary(pr.Summary)
// 	productPhone.SetDescription(pr.Description)
// 	productPhone.SetImageFile(pr.ImageFile)
// 	productPhone.SetPrice(pr.Price)
// 	productPhone.SetProductCode("p" + utils.RandCode(9))
// 	productPhone.SetSKU("ABC-XXX-YYY")
// 	productPhone.SetCreatedAt(time.Now().String())
// 	productPhone.SetUpdatedAt(productPhone.GetCreatedAt())

// 	/*** ObjectID: Bson generate object id ***/
// 	productPhone.SetID(primitive.NewObjectID())

// 	_, err := p.collection.InsertOne(p.ctx, productPhone)

// 	if err != nil {
// 		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
// 			return nil, errors.New("product with that pcode already exists")
// 		}
// 		return nil, err
// 	}
// 	// Create Indexesfor pcode, it help you easy to find product by pcode
// 	opt := options.Index()
// 	opt.SetUnique(true)

// 	index := mongo.IndexModel{Keys: bson.M{"pcode": 1}, Options: opt}

// 	if _, err := p.collection.Indexes().CreateOne(p.ctx, index); err != nil {
// 		return nil, errors.New("could not create index for pcode")
// 	}

// 	var showProduct *entities.Product
// 	// query := bson.M{"_id": res.InsertedID}

// 	query := bson.M{"_id": productPhone.GetID()}

// 	if err = p.collection.FindOne(p.ctx, query).Decode(&showProduct); err != nil {
// 		return nil, err
// 	}

// 	return product, nil
// }

func (p *ProductRepository) FindAllProducts(page int, limit int, currency string) (interface{}, error) {

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

	// calculate exchange rate between base: Euro and dest: currency
	// rate, err := p.getRate(currency)
	// if err != nil {
	// 	p.log.Error("Unable to get rate", "currency", currency, "error", err)
	// }

	// // create a array to contain the rate products
	// pr := entities.Product{}
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

func (p *ProductRepository) FindProductByID(id string, currency string) (*entities.Product, error) {
	// convert string id to objectID
	obId, _ := primitive.ObjectIDFromHex(id)

	// create a query command by id
	query := bson.M{"_id": obId}

	// create container
	var product *entities.Product

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
func (p *ProductRepository) FindProductByName(name string, currency string) ([]*entities.Product, error) {

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

func (p *ProductRepository) FindProductByCategory(category string, currency string) ([]*entities.Product, error) {

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

func (p *ProductRepository) UpdateProduct(id string, pr *models.RequestUpdateProduct) (*entities.Product, error) {

	doc, err := utils.ToDoc(pr)
	if err != nil {
		return nil, err
	}

	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.D{{Key: "_id", Value: obId}}
	update := bson.D{{Key: "$set", Value: doc}}
	res := p.collection.FindOneAndUpdate(p.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatedPost *entities.Product

	if err := res.Decode(&updatedPost); err != nil {
		return nil, errors.New("no post with that Id exists")
	}

	return updatedPost, nil
}
func (p *ProductRepository) DeleteProduct(id string) error {

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
