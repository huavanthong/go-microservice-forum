package worker

import (
	"context"
	"log"
	"time"
)

// OrderWorker defines an order worker.
type OrderWorker struct {
	// add fields here as needed
}

// NewOrderWorker creates a new instance of OrderWorker.
func NewOrderWorker() *OrderWorker {
	return &OrderWorker{}
}

// Run starts the worker.
func (w *OrderWorker) Run(ctx context.Context) error {
	// add logic here to fetch and process orders from the database
	log.Println("Order worker started")
	for {
		select {
		case <-ctx.Done():
			log.Println("Order worker stopped")
			return nil
		default:
			// fetch orders and process them
			time.Sleep(1 * time.Second) // for demo only
		}
	}
}
