package handlers

import (
	"net/http"

	"github.com/huavanthong/microservice-golang/product-api/data"
)

// swagger:route GET /products products listProducts
// Return a list of products from the database
// responses:
//	200: productsResponse

// ListAll handles GET requests and returns all current products
func (p *Products) ListAll(rw http.ResponseWriter, r *http.Request) {

	p.l.Debug("Get all records")

	// set applicatin type to display data on client side
	rw.Header().Add("Content-Type", "application/json")

	// get a currency value based on a query command existed in URL on request
	cur := r.URL.Query().Get("currency")

	// get products with the currency
	prods, err := p.productDB.GetProducts(cur)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	// encode products data to json
	err = data.ToJSON(prods, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Error("Unable to serializing product", "error", err)
	}
}

// swagger:route GET /products/{id} products listSingle
// Return a list of products from the database
// responses:
//	200: productResponse
//	404: errorResponse

// ListSingle handles GET requests
func (p *Products) ListSingle(rw http.ResponseWriter, r *http.Request) {

	// set applicatin type to display data on client side
	rw.Header().Add("Content-Type", "application/json")

	// get id from request
	id := getProductID(r)

	p.l.Debug("Get record", "id", id)

	// get a currency value based on a query command existed in URL on request
	cur := r.URL.Query().Get("currency")

	// find product by id with currency
	prod, err := p.productDB.GetProductByID(id, cur)

	switch err {
	case nil:

	case data.ErrProductNotFound:
		p.l.Error("Unable to fetch product", "error", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	default:
		p.l.Error("Unable to fetching product", "error", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	// convert message response to json
	err = data.ToJSON(prod, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Error("Unable to serializing product", err)
	}
}
