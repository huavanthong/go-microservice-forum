package main

import (
	"log"
	"time"

	"github.com/yourusername/project/internal/workers"
)

func main() {
	// Create a new instance of the OrderWorker.
	orderWorker := workers.NewOrderWorker()

	// Start the worker.
	orderWorker.Start()

	// Wait for some time to allow the worker to process some orders.
	time.Sleep(5 * time.Second)

	// Stop the worker.
	orderWorker.Stop()

	// Log a message indicating that the worker has been stopped.
	log.Printf("Order worker has been stopped.")
}
