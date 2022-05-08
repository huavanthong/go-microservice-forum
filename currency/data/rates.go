package data

import (
	"github.com/hashicorp/go-hclog"
)

type ExchangeRates struct {
	log   hclog.Logger
	rates map[string]float64
}

func NewExchangeRates(l hclog.Logger) (*ExchangeRates, error) {
	// create a object for exchange rates
	er := &ExchangeRates{log: l, rates: map[string]float64{}}

	// assign a value rates for object
	err := er.getRates()

	return er, err
}
