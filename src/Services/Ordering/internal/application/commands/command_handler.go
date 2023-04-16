package commands

import (
	"context"
	"errors"

	"github.com/huavanthong/microservice-golang/src/Services/Ordering/internal/domain/repository"
)

// CommandHandler is the interface for all command handlers
type CommandHandler interface {
	Handle(ctx context.Context, command interface{}) (interface{}, error)
}

// CommandBus is the interface for all command buses
type CommandBus interface {
	RegisterHandler(commandName string, handler CommandHandler)
	Dispatch(ctx context.Context, commandName string, command interface{}) (interface{}, error)
}

// CreateOrderCommand represents the data needed to create an order
type CreateOrderCommand struct {
	UserID      string
	TotalAmount float64
	Items       []OrderItem
}

// DeleteOrderCommand represents the data needed to delete an order
type DeleteOrderCommand struct {
	ID string
}

// UpdateOrderCommand represents the data needed to update an order
type UpdateOrderCommand struct {
	ID          string
	UserID      string
	TotalAmount float64
	Items       []OrderItem
}

// OrderService defines the interface for the order service
type OrderService interface {
	CreateOrder(ctx context.Context, userID string, totalAmount float64, items []OrderItem) (string, error)
	DeleteOrder(ctx context.Context, id string) error
	UpdateOrder(ctx context.Context, id string, userID string, totalAmount float64, items []OrderItem) error
}

// CreateOrderHandler handles the create order command
type CreateOrderHandler struct {
	orderService OrderService
}

// NewCreateOrderHandler creates a new instance of CreateOrderHandler
func NewCreateOrderHandler(orderService OrderService) *CreateOrderHandler {
	return &CreateOrderHandler{orderService: orderService}
}

// Handle handles the create order command
func (h *CreateOrderHandler) Handle(ctx context.Context, command interface{}) (interface{}, error) {
	// Cast the command to CreateOrderCommand
	c, ok := command.(CreateOrderCommand)
	if !ok {
		return nil, errors.New("invalid command")
	}

	// Call the order service to create the order
	orderID, err := h.orderService.CreateOrder(ctx, c.UserID, c.TotalAmount, c.Items)
	if err != nil {
		return nil, err
	}

	// Return the order ID
	return orderID, nil
}

// DeleteOrderHandler handles the delete order command
type DeleteOrderHandler struct {
	orderService OrderService
}

// NewDeleteOrderHandler creates a new instance of DeleteOrderHandler
func NewDeleteOrderHandler(orderService OrderService) *DeleteOrderHandler {
	return &DeleteOrderHandler{orderService: orderService}
}

// Handle handles the delete order command
func (h *DeleteOrderHandler) Handle(ctx context.Context, command interface{}) (interface{}, error) {
	// Cast the command to DeleteOrderCommand
	c, ok := command.(DeleteOrderCommand)
	if !ok {
		return nil, errors.New("invalid command")
	}

	// Call the order service to delete the order
	err := h.orderService.DeleteOrder(ctx, c.ID)
	if err != nil {
		return nil, err
	}

	// Return nil
	return nil, nil
}

// UpdateOrderHandler handles the update order command
type UpdateOrderHandler struct {
	orderService OrderService
}

// NewUpdateOrderHandler creates a new instance of UpdateOrderHandler
func NewUpdateOrderHandler(orderService Order
