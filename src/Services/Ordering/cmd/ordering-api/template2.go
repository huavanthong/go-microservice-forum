package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"

	"github.com/myproject/order/commandservice"
	"github.com/myproject/order/eventhandler"
	"github.com/myproject/order/queryservice"
	"github.com/myproject/order/repository"
	"github.com/myproject/order/service"
)

var mapper = mapper.New()

func init() {
	mapper.CreateMap(&Source{}, &Destination{})
}

func main() {
	// Create a new router instance
	r := mux.NewRouter()

	// Initialize the repository
	orderRepository := repository.NewOrderRepository()

	// Initialize the command service
	orderCommandService := commandservice.NewOrderCommandService(orderRepository)

	// Initialize the query service
	orderQueryService := queryservice.NewOrderQueryService(orderRepository)

	// Initialize the event handler
	orderEventHandler := eventhandler.NewOrderEventHandler()

	// Initialize the order service
	orderService := service.NewOrderService(orderCommandService, orderQueryService, orderEventHandler)

	// Register the JWT middleware before defining the API endpoints
	r.Use(jwtMiddleware)

	// Define the API endpoints
	r.HandleFunc("/orders", orderService.CreateOrderHandler).Methods(http.MethodPost)
	r.HandleFunc("/orders/{id}", orderService.GetOrderHandler).Methods(http.MethodGet)

	// Create a new HTTP server
	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
	}

	// Start the HTTP server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Create a channel to handle OS signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Wait for the OS signal
	<-sigChan

	// Gracefully shutdown the server
	if err := srv.Shutdown(nil); err != nil {
		log.Printf("Failed to shutdown server: %v", err)
	}

	fmt.Println("Server stopped.")
}
