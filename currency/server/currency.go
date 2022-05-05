package server

import (
	"context"

	"github.com/hashicorp/go-hclog"
	protos "github.com/huavanthong/microservice-golang/currency/proto/currency"
)

type Currency struct {
	log hclog.Logger
}

// NewCurrency creates a new Currency server
func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{l}
}

// GetRate
func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle request for GetRate", "base", rr.GetBase(), "dest", rr.GetDestination())
	return &protos.RateResponse{Rate: 0.5}, nil
}
