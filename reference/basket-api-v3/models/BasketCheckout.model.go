package models

type Brand struct {
	UserName      int
	TotalPrice    string
	FirstName     string
	EmailAddress  string
	AddressLine   string
	Country       string
	State         string
	ZipCode       string
	CardName      string
	CardNumber    string
	Expiration    string
	CVV           string
	PaymentMethod string
	CreatedAt     string `json:"created_at" bson:"created_at"`
	UpdatedAt     string `json:"updated_at" bson:"updated_at"`
	DeleteAt      string `json:"deleted_at" bson:"deleted_at"`
}
