package server

import (
	"context"
	"io"
	"time"

	"github.com/hashicorp/go-hclog"
	"github.com/huavanthong/microservice-golang/currency/data"
	protos "github.com/huavanthong/microservice-golang/currency/proto/currency"
)

type Currency struct {
	protos.UnimplementedCurrencyServer
	rates *data.ExchangeRates
	log   hclog.Logger
}

// NewCurrency creates a new Currency server
func NewCurrency(r *data.ExchangeRates, l hclog.Logger) *Currency {
	return &Currency{rates: r, log: l}
}

// GetRate implements the CurrencyServer GetRate method and returns the currency exchange rate
// for the two given currencies.
func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle request for GetRate", "base", rr.GetBase(), "dest", rr.GetDestination())

	// get rates from data package
	rate, err := c.rates.GetRate(rr.GetBase().String(), rr.GetDestination().String())
	if err != nil {
		return nil, err
	}

	return &protos.RateResponse{Rate: rate}, nil
}

func (c *Currency) SubscribeRates(src protos.Currency_SubscribeRatesServer) error {

	// handle client messages
	go func() {
		for {
			rr, err := src.Recv() // Recv is a blocking method which returns on client data
			// io.EOF signals that the client has closed the connection
			if err == io.EOF {
				c.log.Info("Client has closed connection")
				break
			}

			// any other error means the transport between the server and client is unavailable
			if err != nil {
				c.log.Error("Unable to read from client", "error", err)
				break
			}

			c.log.Info("Handle client request", "request_base", rr.GetBase(), "request_dest", rr.GetDestination())
		}
	}()

	// handle server responses
	// we block here to keep the connection open
	for {
		// send a message back to the client
		err := src.Send(&protos.RateResponse{Rate: 12.1})
		if err != nil {
			return err
		}

		time.Sleep(5 * time.Second)
	}
}
