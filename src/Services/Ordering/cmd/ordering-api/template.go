package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/huavanthong/microservice-golang/src/Services/Ordering/internal/application/command_handlers"
	"github.com/huavanthong/microservice-golang/src/Services/Ordering/internal/infrastructure/eventbus"
	"github.com/huavanthong/microservice-golang/src/Services/Ordering/internal/infrastructure/persistence"
)

func main() {
	// Initialize dependencies
	orderRepo := persistence.NewOrderRepository()
	eventBus := eventbus.NewInMemoryEventBus()

	// Initialize command handlers
	createOrderHandler := command_handlers.NewCreateOrderHandler(orderRepo, eventBus)
	deleteOrderHandler := command_handlers.NewDeleteOrderHandler(orderRepo, eventBus)
	updateOrderHandler := command_handlers.NewUpdateOrderHandler(orderRepo, eventBus)

	// Initialize router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/orders", createOrderHandler.Handle).Methods("POST")
	router.HandleFunc("/orders/{id}", deleteOrderHandler.Handle).Methods("DELETE")
	router.HandleFunc("/orders/{id}", updateOrderHandler.Handle).Methods("PUT")

	// Start server
	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
