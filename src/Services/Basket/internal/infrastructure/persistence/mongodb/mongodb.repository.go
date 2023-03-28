package mongodb

import (
	"fmt"

	"github.com/mongodb/mongo-go-driver/bson"

	"github.com/example/shoppingcart/entities"
)

type MongoDBBasketRepository struct {
	session    *mgo.Session
	database   string
	collection string
}

func NewMongoDBBasketRepository(session *mgo.Session, database string, collection string) *MongoDBBasketRepository {
	return &MongoDBBasketRepository{
		session:    session,
		database:   database,
		collection: collection,
	}
}

func (r *MongoDBBasketRepository) Session() *mgo.Session {
	return r.session.Copy()
}

func (r *MongoDBBasketRepository) GetBasket(userName string) (*entities.ShoppingCart, error) {
	session := r.Session()
	defer session.Close()

	c := session.DB(r.database).C(r.collection)

	basket := &entities.ShoppingCart{}
	err := c.Find(bson.M{"username": userName}).One(basket)

	if err != nil {
		return nil, fmt.Errorf("could not get basket for user %s: %v", userName, err)
	}

	return basket, nil
}

func (r *MongoDBBasketRepository) UpdateBasket(basket *entities.ShoppingCart) (*entities.ShoppingCart, error) {
	session := r.Session()
	defer session.Close()

	c := session.DB(r.database).C(r.collection)

	err := c.Update(bson.M{"username": basket.Username}, basket)
	if err != nil {
		return nil, fmt.Errorf("could not update basket for user %s: %v", basket.Username, err)
	}

	return basket, nil
}

func (r *MongoDBBasketRepository) DeleteBasket(userName string) error {
	session := r.Session()
	defer session.Close()

	c := session.DB(r.database).C(r.collection)

	err := c.Remove(bson.M{"username": userName})
	if err != nil {
		return fmt.Errorf("could not delete basket for user %s: %v", userName, err)
	}

	return nil
}
