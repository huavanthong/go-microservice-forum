package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/your-username/order-microservice/internal/api/httpserver"
)

func main() {
	server := httpserver.NewServer()

	// Kết nối router và middleware
	router := httpserver.NewRouter(server)

	// Khởi động server
	addr := fmt.Sprintf(":%d", server.Port)
	log.Printf("Server is running at %s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal(err)
	}
}
