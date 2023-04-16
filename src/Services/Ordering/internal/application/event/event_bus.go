package event

import (
	"sync"

	"github.com/huavanthong/microservice-golang/src/Services/Ordering/internal/domain/events"
)

type EventHandlerFunc func(event interface{}) error

type EventBus struct {
	subscribers map[string][]EventHandlerFunc
	mutex       sync.RWMutex
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]EventHandlerFunc),
	}
}

func (bus *EventBus) Subscribe(eventName string, handler EventHandlerFunc) {
	bus.mutex.Lock()
	defer bus.mutex.Unlock()

	if _, ok := bus.subscribers[eventName]; !ok {
		bus.subscribers[eventName] = []EventHandlerFunc{handler}
	} else {
		bus.subscribers[eventName] = append(bus.subscribers[eventName], handler)
	}
}

func (bus *EventBus) Publish(event interface{}) error {
	bus.mutex.RLock()
	defer bus.mutex.RUnlock()

	eventName := events.GetEventName(event)
	if handlers, ok := bus.subscribers[eventName]; ok {
		for _, handler := range handlers {
			if err := handler(event); err != nil {
				return err
			}
		}
	}
	return nil
}
