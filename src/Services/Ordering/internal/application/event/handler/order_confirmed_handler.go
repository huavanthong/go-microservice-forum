package handler

import (
	"context"
	"errors"
)

type OTPConfirmedEvent struct {
	CustomerID string
}

type OTPConfirmedEventHandler struct {
	// dependencies
	orderService OrderService
}

func (h *OTPConfirmedEventHandler) Handle(ctx context.Context, event OTPConfirmedEvent) error {
	// validate input
	if event.CustomerID == "" {
		return errors.New("invalid input")
	}

	// update order status to "completed"
	err := h.orderService.UpdateOrderStatus(ctx, event.CustomerID, "completed")
	if err != nil {
		return err
	}

	return nil
}
