package commands

import (
	"fmt"
	"time"
)

type CreateOrderCommand struct {
	OrderID    string
	CustomerID string
	OrderDate  time.Time
	Total      float64
}

func (c CreateOrderCommand) Validate() error {
	if c.OrderID == "" {
		return fmt.Errorf("order id is required")
	}
	if c.CustomerID == "" {
		return fmt.Errorf("customer id is required")
	}
	if c.Total <= 0 {
		return fmt.Errorf("total must be greater than zero")
	}
	return nil
}
