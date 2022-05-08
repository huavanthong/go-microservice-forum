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

// for multiple Cube values
type Cubes struct {
	CubeData []Cube `xml:"Cube>Cube>Cube"`
}

// define structure adapt value in vxml
type Cube struct {
	Currency string  `xml:"currency, attr`
	Rate     float64 `xml:"rate, attr"`
}
