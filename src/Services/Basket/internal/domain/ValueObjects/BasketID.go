package ValueObjects

import (
	"encoding/json"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BasketID string

func NewBasketID(id string) (BasketID, error) {
	if id == "" {
		return "", errors.New("basket id cannot be empty")
	}

	return BasketID(id), nil
}

func (id BasketID) Value() (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(string(id))
}

func (id BasketID) String() string {
	return string(id)
}

func (id BasketID) Equals(otherID BasketID) bool {
	return id == otherID
}

func (id BasketID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id)
}

func (id *BasketID) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	basketID, err := NewBasketID(str)
	if err != nil {
		return err
	}

	*id = basketID
	return nil
}
