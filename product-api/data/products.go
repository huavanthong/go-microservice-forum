package data

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/go-hclog"
	protos "github.com/huavanthong/microservice-golang/currency/proto/currency"
)

/************************ Define structure product ************************/

// ErrProductNotFound is an error raised when a product can not be found in the database
var ErrProductNotFound = fmt.Errorf("Product not found")

// Product defines the structure for an API product
// swagger:model
type Product struct {
	// the id for the product
	//
	// required: false
	// min: 1
	ID int `json:"id"` // Unique identifier for the product

	// the name for this poduct
	//
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required,gt=0"`

	// the description for this poduct
	//
	// required: false
	// max length: 10000
	Description string `json:"description"`

	// the price for the product
	//
	// required: true
	// min: 0.01
	Price float64 `json:"price" validate:"required,gt=0"`

	// the SKU for the product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU string `json:"sku" validate:"sku"`

	CreatedOn string `json:"-"`
	UpdatedOn string `json:"-"`
	DeletedOn string `json:"-"`
}

// Products defines a slice of Product
type Products []*Product

// ProductsDB: defines a DB of Product with the currency service
type ProductsDB struct {
	currency protos.CurrencyClient                // For registration a currency service
	log      hclog.Logger                         // For logger
	rates    map[string]float64                   // For a specific exchanges rates
	client   protos.Currency_SubscribeRatesClient // For client want to subscribe a interval updated exchanges rates
}

func NewProductsDB(c protos.CurrencyClient, l hclog.Logger) *ProductsDB {
	pb := &ProductsDB{c, l, make(map[string]float64), nil}

	go pb.handleUpdates()

	return pb
}

// Implement handler to update exchange rates from the currency service.
func (p *ProductsDB) handleUpdates() {
	// call subscriber handler from the currency service.
	sub, err := p.currency.SubscribeRates(context.Background())
	if err != nil {
		p.log.Error("Unable to subscribe for rates", "error", err)
	}

	// assign subscriber handler to client.
	// right now, client can follow any changes on currency
	p.client = sub

	// wait updating
	for {
		rr, err := sub.Recv()
		p.log.Info("Recieved updated rate from server", "dest", rr.GetDestination().String())

		if err != nil {
			p.log.Error("Error receiving message", "error", err)
			return
		}

		p.rates[rr.Destination.String()] = rr.Rate
	}
}

/************************ Method for Product ************************/
/************ GET ************/
// GetProducts returns a list of products
func (p *ProductsDB) GetProducts(currency string) (Products, error) {

	// if currency is empty, it return productList with the default of
	// base currency
	if currency == "" {
		return productList, nil
	}

	// calculate exchange rate between base: Euro and dest: currency
	rate, err := p.getRate(currency)
	if err != nil {
		p.log.Error("Unable to get rate", "currency", currency, "error", err)
	}

	// create a array to contain the rate products
	pr := Products{}
	// loop in productList to update to the product with rate
	for _, p := range productList {
		// get a product
		np := *p
		// update it's currency with rate
		np.Price = np.Price * rate
		// push to a temp storage of product
		pr = append(pr, &np)
	}
	return pr, nil
}

// GetProductByID returns a single product which matches the id from the
// database.
// If a product is not found this function returns a ProductNotFound error
func (p *ProductsDB) GetProductByID(id int, currency string) (*Product, error) {
	// find product by id
	i := findIndexByProductID(id)
	if id == -1 {
		return nil, ErrProductNotFound
	}

	// if currency is empty, it return productList with the default of
	// base currency
	if currency == "" {
		return productList[i], nil
	}

	// calculate exchange rate between base: Euro and dest: currency
	rate, err := p.getRate(currency)
	if err != nil {
		p.log.Error("Unable to get rate", "currency", currency, "error", err)
	}

	// get product in list
	np := *productList[i]
	// update it's currency with rate
	np.Price = np.Price * rate

	return &np, nil
}

/************ POST ************/
// AddProduct addies a product to list
func (p *ProductsDB) AddProduct(pr *Product) {
	pr.ID = getNextID()
	productList = append(productList, pr)
}

/************ PUT ************/
// UpdateProduct updates info to product
func (p *ProductsDB) UpdateProduct(pr Product) error {

	i := findIndexByProductID(pr.ID)
	if i == -1 {
		return ErrProductNotFound
	}

	// update the product in the DB
	productList[i] = &pr

	return nil
}

/************ DELETE ************/
// DeleteProduct deletes a product from the database
func (p *ProductsDB) DeleteProduct(id int) error {
	i := findIndexByProductID(id)
	if i == -1 {
		return ErrProductNotFound
	}

	productList = append(productList[:i], productList[i+1])

	return nil
}

/************************ Internal function for Product ************************/
func getNextID() int {
	// get ID at the last product in productList
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

func findProduct(id int) (*Product, int, error) {

	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

// findIndex finds the index of a product in the database
// returns -1 when no product can be found
func findIndexByProductID(id int) int {
	for i, p := range productList {
		if p.ID == id {
			return i
		}
	}

	return -1
}

func (p *ProductsDB) getRate(destination string) (float64, error) {
	// define base currency is Euro
	rr := &protos.RateRequest{
		Base:        protos.Currencies(protos.Currencies_value["EUR"]),
		Destination: protos.Currencies(protos.Currencies_value[destination]),
	}

	resp, err := p.currency.GetRate(context.Background(), rr)
	return resp.Rate, err
}

/************************ Storage Product ************************/
// productList is a hard coded list of products for this
// example data source
var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
