package data

import (
	"encoding/xml"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/go-hclog"
)

type ExchangeRates struct {
	log   hclog.Logger
	rates map[string]float64
}

func NewRates(l hclog.Logger) (*ExchangeRates, error) {
	// create a object for exchange rates
	er := &ExchangeRates{log: l, rates: map[string]float64{}}

	// assign a value rates for object
	err := er.getRates()

	return er, err
}

// GetRate: expose API to calcuate exchange rate for base and dest
func (e *ExchangeRates) GetRate(base, dest string) (float64, error) {

	// get rate the currency for base
	br, ok := e.rates[base]
	if !ok {
		return 0, fmt.Errorf("Rate not found for currency %s", base)
	}

	// get rate the currency for dest
	dr, ok := e.rates[dest]
	if !ok {
		return 0, fmt.Errorf("Rate not found for currency %s", dest)
	}

	return dr / br, nil
}

// MonitorRates checks the rates in the ECB API every interval and sends a message to the
// returned channel when there are changes
//
// Note: the ECB API only returns data once a day, this function only simulates the changes
// in rates for demonstration purposes
func (e *ExchangeRates) MonitorRates(interval time.Duration) chan struct{} {
	ret := make(chan struct{})

	// We use goroutine because we want to wait it forever
	go func() {
		// Register time with a period
		ticker := time.NewTicker(interval)
		for {
			select {
			case <-ticker.C:
				// just add a random difference to the rate and return it
				// this simulates the fluctuations in currency rates
				for k, v := range e.rates {
					// change can be 10% of original value
					change := (rand.Float64() / 10)
					// is this a postive or negative change
					direction := rand.Intn(1)

					if direction == 0 {
						// new value with be min 90% of old
						change = 1 - change
					} else {
						// new value will be 110% of old
						change = 1 + change
					}

					// modify the rate
					e.rates[k] = v * change
				}

				// notify updates, this will block unless there is a listener on the other end
				ret <- struct{}{}
			}
		}
	}()

	return ret
}

// define internal value to getRates
func (e *ExchangeRates) getRates() error {
	// using http to get the exchange rate of currency from european central bank
	resp, err := http.DefaultClient.Get("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml")
	if err != nil {
		return nil
	}

	// handling error code from response using http package
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Expected error code 200 got %d", resp.StatusCode)
	}
	// any abnormal case, it need to be closed
	defer resp.Body.Close()

	// initialize a container to collect data
	md := &Cubes{}

	// create a decoder, and decode xml to get a raw data
	xml.NewDecoder(resp.Body).Decode(&md)

	// loop to parse all values to from the container of data
	// note: for each currency in CubeData
	for _, c := range md.CubeData {
		// parse rate value with 64 bit, from c.Rate
		r, err := strconv.ParseFloat(c.Rate, 64)
		if err != nil {
			return err
		}

		// assign data
		e.rates[c.Currency] = r
	}

	e.rates["EUR"] = 1

	return nil
}

// for multiple Cube values
type Cubes struct {
	CubeData []Cube `xml:"Cube>Cube>Cube"`
}

// define structure adapt value in vxml
type Cube struct {
	Currency string `xml:"currency, attr`
	Rate     string `xml:"rate,attr"`
}
