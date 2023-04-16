package handler

import (
	"context"
	"encoding/json"
	"log"

	"github.com/huavanthong/microservice-golang/src/Services/Ordering/internal/domain/events"
	"github.com/huavanthong/microservice-golang/src/Services/Ordering/internal/domain/repository"
)

type OrderCreatedHandler struct {
	orderRepo repository.OrderRepository
}

func NewOrderCreatedHandler(orderRepo repository.OrderRepository) *OrderCreatedHandler {
	return &OrderCreatedHandler{orderRepo: orderRepo}
}

func (h *OrderCreatedHandler) Handle(ctx context.Context, data []byte) error {
	// Parse event data
	event := &events.OrderCreatedEvent{}
	err := json.Unmarshal(data, event)
	if err != nil {
		log.Printf("Error parsing event data: %v", err)
		return err
	}

	// Update order status to "Created" in repository
	err = h.orderRepo.UpdateStatus(ctx, event.OrderID, "Created")
	if err != nil {
		log.Printf("Error updating order status to Created: %v", err)
		return err
	}

	log.Printf("Order status updated to Created for order %s", event.OrderID)

	return nil
}
