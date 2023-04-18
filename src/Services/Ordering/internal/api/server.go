package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

// Server represents the HTTP server for the order microservice.
type Server struct {
	httpServer *http.Server
	router     *mux.Router
}

// NewServer returns a new instance of the HTTP server.
func NewServer(addr string) *Server {
	router := mux.NewRouter()
	server := &Server{
		httpServer: &http.Server{
			Addr:         addr,
			Handler:      router,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		router: router,
	}

	server.initRoutes()

	return server
}

// initRoutes initializes the HTTP routes for the server.
func (s *Server) initRoutes() {
	// Order HTTP handler
	orderHandler := &OrderHTTPHandler{}

	// API endpoints
	s.router.HandleFunc("/orders", orderHandler.CreateOrder).Methods(http.MethodPost)
	s.router.HandleFunc("/orders/{id}", orderHandler.GetOrder).Methods(http.MethodGet)
	s.router.HandleFunc("/orders/{id}", orderHandler.UpdateOrder).Methods(http.MethodPut)
	s.router.HandleFunc("/orders/{id}", orderHandler.DeleteOrder).Methods(http.MethodDelete)

	// Health check endpoint
	s.router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
}

// Start starts the HTTP server.
func (s *Server) Start() {
	log.Printf("Starting server on %s\n", s.httpServer.Addr)

	if err := s.httpServer.ListenAndServe(); err != nil {
		log.Fatal(fmt.Sprintf("Server failed to start: %v", err))
		os.Exit(1)
	}
}

// Stop stops the HTTP server.
func (s *Server) Stop() error {
	log.Printf("Stopping server on %s\n", s.httpServer.Addr)
	return s.httpServer.Shutdown(nil)
}
