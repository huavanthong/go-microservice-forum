package main

import (
	"log"
	"net/http"

	"your-app/internal/api/httpserver"
)

func main() {
	// Tạo một router và gán các route và middleware cho nó
	router1 := httpserver.NewRouter()
	router1.Use(httpserver.LoggingMiddleware)
	router1.Use(httpserver.AuthenticationMiddleware)

	// Tạo một server với cổng 8080 và router1
	server1 := &http.Server{
		Addr:    ":8080",
		Handler: router1,
	}

	// Tạo một router khác và gán các route và middleware khác cho nó
	router2 := httpserver.NewRouter()
	router2.Use(httpserver.CORSMiddleware)
	router2.Use(httpserver.AuthenticationMiddleware)

	// Tạo một server khác với cổng 8081 và router2
	server2 := &http.Server{
		Addr:    ":8081",
		Handler: router2,
	}

	// Khởi động các server
	go func() {
		if err := server1.ListenAndServe(); err != nil {
			log.Fatalf("Server 1 failed to start: %v", err)
		}
	}()

	go func() {
		if err := server2.ListenAndServe(); err != nil {
			log.Fatalf("Server 2 failed to start: %v", err)
		}
	}()

	// Chờ tín hiệu để dừng server
	// ...

}
