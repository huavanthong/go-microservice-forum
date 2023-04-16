package event

import "time"

type OrderCreated struct {
	ID        string
	Customer  string
	Total     float64
	Items     []OrderItem
	CreatedAt time.Time
}

type OrderItem struct {
	ProductID string
	Quantity  int
	Price     float64
}
