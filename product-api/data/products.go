package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator"
)

/************************ Define structure product ************************/
// Product defines the structure for an API product
type Product struct {
	ID          int     `json:"id`
	Name        string  `json:"name" validate:"require"` // Validate-Step1: require to validate for this member.
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"require,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// Products is a collection of Product
type Products []*Product

/************************ Validation ************************/
func (p *Product) Validate() error {
	// Validate-Step2: create a object for this validator
	validate := validator.New()

	// Validate-Step3: register a customer validator function
	validate.RegisterValidation("sku", validateSKU)

	// Validate-Step4: validate structure product
	return validate.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	// sku is of format abc-absd-dfsdf
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}

	return true
}

/************************ Method for Product ************************/
// GetProducts returns a list of products
func GetProducts() Products {
	return productList
}

// AddProduct addies a product to list
func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

// UpdateProduct updates info to product
func UpdateProduct(id int, p *Product) error {

	// find pos in productList from id
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	// update info product
	p.ID = id
	productList[pos] = p

	return nil
}

/************************ Internal function for Product ************************/
func getNextID() int {
	// get ID at the last product in productList
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {

	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
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
