package entities

type BasketCheckout struct {
	UserName      string  `json:"userName"`
	TotalPrice    float64 `json:"totalPrice"`
	FirstName     string  `json:"firstName"`
	LastName      string  `json:"lastName"`
	EmailAddress  string  `json:"emailAddress"`
	AddressLine   string  `json:"addressLine"`
	Country       string  `json:"country"`
	State         string  `json:"state"`
	ZipCode       string  `json:"zipCode"`
	CardName      string  `json:"cardName"`
	CardNumber    string  `json:"cardNumber"`
	Expiration    string  `json:"expiration"`
	CVV           string  `json:"cvv"`
	PaymentMethod int     `json:"paymentMethod"`
}
