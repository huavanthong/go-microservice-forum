package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/huavanthong/microservice-golang/src/Services/Ordering/internal/api/httpserver"

	"github.com/huavanthong/microservice-golang/src/Services/Ordering/internal/application/commands"
	"github.com/huavanthong/microservice-golang/src/Services/Ordering/internal/application/queries"
)

func main() {

	// Initialize command bus
	commandBus := commands.NewCommandBus()
	query := queries.NewQuery()

	httpHandler := httpserver.NewHttpHandler(commandBus, query)

	// Kết nối router và middleware
	router := httpserver.NewRouter(httpHandler)

	server := httpserver.NewServer("8004")

	// Khởi động server
	addr := fmt.Sprintf(":%d", server.Port)
	log.Printf("Server is running at %s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal(err)
	}
}
