package worker

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Config defines the configuration for the worker.
type Config struct {
	// add configuration fields here as needed
}

// Worker defines the interface for a worker.
type Worker interface {
	Run(context.Context) error
}

// NewWorker creates a new instance of the worker.
func NewWorker(config *Config, worker Worker) *worker {
	return &worker{
		config: config,
		worker: worker,
	}
}

// worker represents a worker.
type worker struct {
	config *Config
	worker Worker
}

// Run starts the worker.
func (w *worker) Run() error {
	// set up context and cancel function for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// set up signal handler to trigger shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-sig
		log.Println("Shutting down worker...")
		cancel()
	}()

	// start the worker
	log.Println("Starting worker...")
	if err := w.worker.Run(ctx); err != nil {
		log.Printf("Worker stopped with error: %v\n", err)
		return err
	}

	return nil
}
