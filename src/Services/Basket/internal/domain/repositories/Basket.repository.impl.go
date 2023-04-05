package repositories

import (
	"fmt"

	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/entities"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/interfaces/persistence"
)

type BasketRepositoryImpl struct {
	mongoPersistence persistence.BasketPersistence
	redisPersistence persistence.BasketPersistence
}

func NewBasketRepositoryImpl(mongoPersistence persistence.BasketPersistence, redisPersistence persistence.BasketPersistence) BasketRepository {
	return &BasketRepositoryImpl{mongoPersistence, redisPersistence}
}

func (br *BasketRepositoryImpl) CreateBasket(userId string) (*entities.Basket, error) {

	// Create basket in Redis
	fmt.Println("Check 1: ")
	basket, err := br.redisPersistence.Create(userId)
	if err != nil {
		return nil, err
	}
	fmt.Println("Check 2: ", err)
	fmt.Println("Check 2: ", basket)
	// Create basket in MongoDB
	_, err = br.mongoPersistence.Create(userId)
	if err != nil {
		// Rollback Redis basket creation on error
		br.redisPersistence.Delete(userId)
		return nil, err
	}
	fmt.Println("Check 3: ", err)

	return basket, nil
}

func (br *BasketRepositoryImpl) GetBasket(userName string) (*entities.Basket, error) {
	// Try to get basket from Redis
	basket, err := br.redisPersistence.GetByUserName(userName)
	if err != nil {
		// Try to get basket from MongoDB
		basket, err = br.mongoPersistence.GetByUserName(userName)
		if err != nil {
			return nil, err
		}

		// Cache basket in Redis
		basket, err = br.redisPersistence.Update(basket)
		if err != nil {
			return nil, err
		}
	}

	return basket, nil
}

func (br *BasketRepositoryImpl) UpdateBasket(userName string, cart *entities.Basket) (*entities.Basket, error) {
	// Update basket in Redis
	if _, err := br.redisPersistence.Update(cart); err != nil {
		return nil, err
	}

	// Update basket in MongoDB
	if _, err := br.mongoPersistence.Update(cart); err != nil {
		// Rollback Redis basket update on error
		oldCart, err := br.redisPersistence.GetByUserName(userName)
		if err != nil {
			return nil, err
		}
		if _, err := br.redisPersistence.Update(oldCart); err != nil {
			return nil, err
		}

		return nil, err
	}

	return cart, nil
}

func (br *BasketRepositoryImpl) DeleteBasket(userName string) error {
	// Delete basket from Redis
	if err := br.redisPersistence.Delete(userName); err != nil {
		return err
	}

	// Delete basket from MongoDB
	if err := br.mongoPersistence.Delete(userName); err != nil {
		// Rollback Redis basket deletion on error
		oldCart, err := br.redisPersistence.Create(userName)
		if err != nil {
			return err
		}
		if _, err := br.redisPersistence.Update(oldCart); err != nil {
			return err
		}

		return err
	}

	return nil
}
