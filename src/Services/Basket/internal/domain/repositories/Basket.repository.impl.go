package repositories

import (
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/entities"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/interfaces/api/models"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/interfaces/persistence"
)

type BasketRepositoryImpl struct {
	mongoPersistence persistence.BasketPersistence
	redisPersistence persistence.BasketPersistence
}

func NewBasketRepositoryImpl(mongoPersistence persistence.BasketPersistence, redisPersistence persistence.BasketPersistence) BasketRepository {
	return &BasketRepositoryImpl{mongoPersistence, redisPersistence}
}

func (br *BasketRepositoryImpl) CreateBasket(cbr *models.CreateBasketRequest) (*entities.Basket, error) {

	userId := cbr.UserID

	// Try to get basket from MongoDB
	basket, err := br.mongoPersistence.Get(userId)
	if err != nil {
		// Try to get basket from Redis
		basket, err = br.redisPersistence.Get(userId)
		if err != nil {
			// Basket not found in MongoDB and Redis, create a new one
			basket, err = br.redisPersistence.Create(userId)
			if err != nil {
				return nil, err
			}

			// Create basket in MongoDB
			_, err = br.mongoPersistence.Create(userId)
			if err != nil {
				// Rollback Redis basket creation on error
				br.redisPersistence.Delete(userId)
				return nil, err
			}

			// Generate info response
			basket.UserName = cbr.UserName
		}
	}

	return basket, nil
}

func (br *BasketRepositoryImpl) GetBasket(userId string) (*entities.Basket, error) {

	// Try to get basket from Redis
	basket, err := br.redisPersistence.Get(userId)
	if err != nil {
		// Try to get basket from MongoDB
		basket, err = br.mongoPersistence.Get(userId)
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

func (br *BasketRepositoryImpl) UpdateBasket(basket *entities.Basket) (*entities.Basket, error) {

	// Update basket in Redis
	if _, err := br.redisPersistence.Update(basket); err != nil {
		return nil, err
	}

	// Update basket in MongoDB
	if _, err := br.mongoPersistence.Update(basket); err != nil {
		// Rollback Redis basket update on error
		oldCart, err := br.redisPersistence.Get(basket.UserID)
		if err != nil {
			return nil, err
		}
		if _, err := br.redisPersistence.Update(oldCart); err != nil {
			return nil, err
		}

		return nil, err
	}

	return basket, nil
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
