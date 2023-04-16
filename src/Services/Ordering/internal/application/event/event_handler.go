package event

import (
	"context"
	"encoding/json"

	"github.com/huavanthong/microservice-golang/src/Services/Ordering/internal/application/commands"
	"github.com/huavanthong/microservice-golang/src/Services/Ordering/internal/application/query"
	"github.com/myusername/order-microservice/internal/app/repository"
)

type EventHandler struct {
	commandBus commands.Bus
	queryBus   query.Bus
	repo       repository.Repository
}

func NewEventHandler(commandBus commands.Bus, queryBus query.Bus, repo repository.Repository) *EventHandler {
	return &EventHandler{
		commandBus: commandBus,
		queryBus:   queryBus,
		repo:       repo,
	}
}

func (h *EventHandler) Handle(ctx context.Context, event interface{}) error {
	switch e := event.(type) {
	case OrderCreated:
		// Handle OrderCreated event
		orderData, err := json.Marshal(e.OrderData)
		if err != nil {
			return err
		}
		cmd := command.CreateOrder{
			UserID:    e.UserID,
			OrderData: orderData,
		}
		return h.commandBus.Send(ctx, &cmd)

	case OrderUpdated:
		// Handle OrderUpdated event
		orderData, err := json.Marshal(e.OrderData)
		if err != nil {
			return err
		}
		cmd := command.UpdateOrder{
			OrderID:   e.OrderID,
			UserID:    e.UserID,
			OrderData: orderData,
		}
		return h.commandBus.Send(ctx, &cmd)

	case OrderDeleted:
		// Handle OrderDeleted event
		cmd := command.DeleteOrder{
			OrderID: e.OrderID,
			UserID:  e.UserID,
		}
		return h.commandBus.Send(ctx, &cmd)

	case PaymentProcessed:
		// Handle PaymentProcessed event
		order, err := h.repo.GetOrder(ctx, e.OrderID, e.UserID)
		if err != nil {
			return err
		}
		if order.PaymentStatus != "paid" {
			order.PaymentStatus = "paid"
			return h.repo.UpdateOrder(ctx, order)
		}

	case OrderShipped:
		// Handle OrderShipped event
		order, err := h.repo.GetOrder(ctx, e.OrderID, e.UserID)
		if err != nil {
			return err
		}
		if order.ShippingStatus != "shipped" {
			order.ShippingStatus = "shipped"
			return h.repo.UpdateOrder(ctx, order)
		}
	}

	return nil
}
