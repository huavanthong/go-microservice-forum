package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/huavanthong/microservice-golang/product-api/handlers"
	"github.com/nicholasjackson/env"
)

// declare environment server
var bindAddress = env.String("BIND_ADDRESS", false, ":8080", "Bind address for the server")

func main() {

	env.Parse()

	// create logger only for product-api
	l := log.New(os.Stdout, "products-api", log.LstdFlags)

	// create the handlers
	ph := handlers.NewProducts(l)

	// create a new server mux and register the handlers
	sm := http.NewServeMux()
	sm.Handle("/", ph)

	// create a new server
	s := http.Server{
		Addr:         *bindAddress,      // configure the bind address
		Handler:      sm,                // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()
}
