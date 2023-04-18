package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	handler *HttpHandler
}

func NewRouter(handler *HttpHandler) *Router {
	return &Router{
		handler: handler,
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router().ServeHTTP(w, req)
}

func (r *Router) router() http.Handler {
	router := mux.NewRouter()

	// Endpoint for creating a new discount
	router.HandleFunc("/discounts", r.handler.CreateDiscount).Methods(http.MethodPost)

	// Endpoint for retrieving a discount by code
	router.HandleFunc("/discounts/{code}", r.handler.GetDiscount).Methods(http.MethodGet)

	return router
}
