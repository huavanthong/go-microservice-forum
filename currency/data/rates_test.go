package data

import (
	"fmt"
	"testing"

	"github.com/hashicorp/go-hclog"
	protos "github.com/huavanthong/microservice-golang/currency/proto/currency"
)

/*********************************************************************
Testing
/*********************************************************************/
/*********************** NewRates ***********************/
// Case 1: Normal case
func TestNewRates(t *testing.T) {
	// NewRates: create a handler ExchangeRates
	tr, err := NewRates(hclog.Default())

	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Rates %#v", tr.rates)
}

/*********************** GetRate ***********************/
// Case 1: Normal case
// 			Get exchange rates from GBP to USD
func TestGetRate(t *testing.T) {
	// NewRates: create a handler ExchangeRates
	tr, err := NewRates(hclog.Default())

	// define base currency is GBP
	rr := &protos.RateRequest{
		Base:        protos.Currencies(protos.Currencies_value["EUR"]),
		Destination: protos.Currencies(protos.Currencies_value["USD"]),
	}

	rate, err := tr.GetRate(rr.GetBase().String(), rr.GetDestination().String())
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("GetRates: from %#v to %#v: %#v", rr.GetBase().String(), rr.GetDestination().String(), rate)
}

// Benchmark 1:
// 			Measure time for GetRates
func BenchmarkGetRate(b *testing.B) {

	// NewRates: create a handler ExchangeRates
	tr, _ := NewRates(hclog.Default())

	// define base currency is GBP
	rr := &protos.RateRequest{
		Base:        protos.Currencies(protos.Currencies_value["EUR"]),
		Destination: protos.Currencies(protos.Currencies_value["USD"]),
	}

	for i := 0; i < b.N; i++ {
		rate, _ := tr.GetRate(rr.GetBase().String(), rr.GetDestination().String())

		fmt.Printf("GetRates: from %#v to %#v: %#v", rr.GetBase().String(), rr.GetDestination().String(), rate)

	}
}
