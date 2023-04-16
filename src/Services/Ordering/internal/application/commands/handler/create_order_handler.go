package handler

import (
	"context"
	"errors"
	"time"

	"github.com/huavanthong/microservice-golang/src/Services/Ordering/internal/domain/entity"
	"github.com/huavanthong/microservice-golang/src/Services/Ordering/internal/domain/repository"
)

type CreateOrderHandler struct {
	orderRepo repository.OrderRepository
}

type CreateOrderRequest struct {
	UserID      string
	TotalAmount float64
	Items       []entity.OrderItem
}

type CreateOrderResponse struct {
	ID string
}

type OrderService interface {
	CreateOrder(ctx context.Context, userID string, totalAmount float64, items []entity.OrderItem) (string, error)
}

func NewCreateOrderHandler(orderRepo repository.OrderRepository) *CreateOrderHandler {
	return &CreateOrderHandler{orderRepo: orderRepo}
}

func (h *CreateOrderHandler) Handle(ctx context.Context, req CreateOrderRequest) (*CreateOrderResponse, error) {
	// Validate request
	if len(req.Items) == 0 {
		return nil, errors.New("invalid request: empty items")
	}

	// Create order entity
	order := &entity.Order{
		UserID:      req.UserID,
		TotalAmount: req.TotalAmount,
		Items:       req.Items,
		CreatedAt:   time.Now(),
	}

	// Save order to repository
	id, err := h.orderRepo.Save(ctx, order)
	if err != nil {
		return nil, err
	}

	// Return response
	return &CreateOrderResponse{ID: id}, nil
}

func (h *CreateOrderHandler) CreateOrder(ctx context.Context, userID string, totalAmount float64, items []domain.OrderItem) (string, error) {
	req := CreateOrderRequest{UserID: userID, TotalAmount: totalAmount, Items: items}
	res, err := h.Handle(ctx, req)
	if err != nil {
		return "", err
	}
	return res.ID, nil
}
