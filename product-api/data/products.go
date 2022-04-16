package data

import (
	"fmt"
	"time"
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

// Products defines a slice of Product
type Products []*Product

/************************ Method for Product ************************/
/************ GET ************/
// GetProducts returns a list of products
func GetProducts() Products {
	return productList
}

// GetProductByID returns a single product which matches the id from the
// database.
// If a product is not found this function returns a ProductNotFound error
func GetProductByID(id int) (*Product, error) {
	i := findIndexByProductID(id)
	if id == -1 {
		return nil, ErrProductNotFound
	}

	return productList[i], nil
}

/************ POST ************/
// AddProduct addies a product to list
func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

/************ PUT ************/
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

/************ DELETE ************/
// DeleteProduct deletes a product from the database
func DeleteProduct(id int) error {
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

var ErrProductNotFound = fmt.Errorf("Product not found")

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
