package event_handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/huavanthong/microservice-golang/src/Services/Ordering/internal/domain/events"
	"github.com/huavanthong/microservice-golang/src/Services/Ordering/internal/domain/repository"
	"github.com/huavanthong/microservice-golang/src/Services/Ordering/internal/service/command_handlers"
)

type EventHandler struct {
	orderRepo  repository.OrderRepository
	commandBus *command_handlers.CommandBus
	eventBus   events.EventBus
}

func NewEventHandler(orderRepo repository.OrderRepository, commandBus *command_handlers.CommandBus, eventBus events.EventBus) *EventHandler {
	handler := &EventHandler{
		orderRepo:  orderRepo,
		commandBus: commandBus,
		eventBus:   eventBus,
	}

	// Register command handlers
	handler.commandBus.RegisterHandler(reflect.TypeOf(command_handlers.CreateOrderCommand{}), command_handlers.NewCreateOrderHandler(orderRepo))

	return handler
}

func (h *EventHandler) HandleEvent(ctx context.Context, event events.Event) error {
	switch event.EventType() {
	case events.OrderCreatedEventType:
		orderCreatedEvent := &events.OrderCreatedEvent{}
		err := json.Unmarshal(event.Payload(), orderCreatedEvent)
		if err != nil {
			return fmt.Errorf("error unmarshalling OrderCreatedEvent: %w", err)
		}

		// Send command to create order
		createOrderCommand := &command_handlers.CreateOrderCommand{
			UserID:      orderCreatedEvent.UserID,
			TotalAmount: orderCreatedEvent.TotalAmount,
			Items:       orderCreatedEvent.Items,
		}
		err = h.commandBus.Dispatch(ctx, createOrderCommand)
		if err != nil {
			return fmt.Errorf("error sending CreateOrderCommand: %w", err)
		}
	}

	return nil
}
