package ValueObjects

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BasketID string

func NewBasketID() BasketID {

	id := string(primitive.NewObjectID().Hex())

	return BasketID(id)
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
