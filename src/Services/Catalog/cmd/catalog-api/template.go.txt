package main

import (
	"log"
	"net/http"
)

func main() {
	// định nghĩa handler cho API endpoint
	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		// trả về danh sách sản phẩm dưới dạng JSON
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"products":[{"id":1,"name":"product 1"},{"id":2,"name":"product 2"}]}`))
	})

	// khởi chạy HTTP server
	log.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
