package value_objects

import (
	"errors"
)

type Stock struct {
	available         int
	restockThreshold  int
	maxStockThreshold int
}

func NewStock(available int, restockThreshold int, maxStockThreshold int) (*Stock, error) {
	if available < 0 || restockThreshold <= 0 || maxStockThreshold <= 0 {
		return nil, errors.New("Invalid argument")
	}

	if available > maxStockThreshold {
		return nil, errors.New("Available stock cannot be greater than max stock threshold.")
	}

	return &Stock{
		available:         available,
		restockThreshold:  restockThreshold,
		maxStockThreshold: maxStockThreshold,
	}, nil
}

func (s *Stock) Available() int {
	return s.available
}

func (s *Stock) RestockThreshold() int {
	return s.restockThreshold
}

func (s *Stock) MaxStockThreshold() int {
	return s.maxStockThreshold
}
