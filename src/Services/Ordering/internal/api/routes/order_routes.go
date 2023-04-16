package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/myapp/order-microservice/handler"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// Define the API endpoints
	r.HandleFunc("/orders", handler.CreateOrderHandler).Methods(http.MethodPost)
	r.HandleFunc("/orders/{id}", handler.GetOrderHandler).Methods(http.MethodGet)

	return r
}
