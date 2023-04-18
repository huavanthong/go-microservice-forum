package httpserver

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

	// Endpoint for creating a new order
	router.HandleFunc("/ordering", r.handler.CreateOrder).Methods(http.MethodPost)

	// Endpoint for retrieving a order by code
	router.HandleFunc("/ordering/{code}", r.handler.GetOrder).Methods(http.MethodGet)

	// Endpoint for updating a order
	router.HandleFunc("/ordering", r.handler.GetOrder).Methods(http.MethodPatch)

	// Endpoint for deleting a order by code
	router.HandleFunc("/ordering/{code}", r.handler.GetOrder).Methods(http.MethodDelete)

	return router
}
