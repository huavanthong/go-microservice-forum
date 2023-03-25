package value_object

import 

type Price struct {
	amount decimal.Decimal
}

func NewPrice(amount float64) (*Price, error) {
	amt := decimal.NewFromFloat(amount)
	if amt.LessThan(decimal.Zero) {
		return nil, errors.New("amount must be greater than zero")
	}
	return &Price{
		amount: amt,
	}, nil
}

func (p *Price) Amount() float64 {
	return p.amount.Float64()
}

func (p *Price) Add(other *Price) *Price {
	return &Price{
		amount: p.amount.Add(other.amount),
	}
}
