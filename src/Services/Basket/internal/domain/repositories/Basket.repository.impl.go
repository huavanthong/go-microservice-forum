package repositories

import (
	"github.com/sirupsen/logrus"

	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/entities"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/interfaces/persistence"
)

type BasketRepositoryImpl struct {
	logger           *logrus.Entry
	mongoPersistence persistence.BasketPersistence
	redisPersistence persistence.BasketPersistence
}

func NewBasketRepositoryImpl(logger *logrus.Entry, mongoPersistence persistence.BasketPersistence, redisPersistence persistence.BasketPersistence) BasketRepository {
	return &BasketRepositoryImpl{
		logger:           logrus.WithField("module", "BasketRepositoryImpl"),
		mongoPersistence: mongoPersistence,
		redisPersistence: redisPersistence,
	}
}

func (br *BasketRepositoryImpl) CreateBasket(basket *entities.Basket) (*entities.Basket, error) {

	var create_basket *entities.Basket
	// Extract user Id from request
	userId := basket.UserID

	// Log the user ID being processed
	br.logger.Info("Creating basket for user ID %s", userId)

	// Try to get basket from MongoDB
	find_basket, err := br.mongoPersistence.Get(userId)
	if err != nil {
		br.logger.Errorf("Failed to get basket from MongoDB: %s", err.Error())

		// Try to get basket from Redis
		find_basket, err = br.redisPersistence.Get(userId)
		if err != nil {
			br.logger.Errorf("Failed to get basket from Redis: %s", err.Error())
			return nil, err
		}

		if find_basket == nil {
			// Basket not found in MongoDB and Redis, create a new one
			create_basket, err = br.redisPersistence.Create(basket)
			if err != nil {
				br.logger.Errorf("Failed to create basket in Redis: %s", err.Error())
				return nil, err
			}

			// Create basket in MongoDB
			create_basket, err = br.mongoPersistence.Create(basket)
			if err != nil {
				// Rollback Redis basket creation on error
				br.redisPersistence.Delete(userId)
				br.logger.Errorf("Failed to create basket in MongoDB: %s", err.Error())
				return nil, err
			}
		}
	}

	// Log the basket being returned
	br.logger.Info("Created basket %+v for user ID %s", create_basket, create_basket.UserID)

	return create_basket, nil
}

func (br *BasketRepositoryImpl) GetBasket(userId string) (*entities.Basket, error) {

	// Log the user ID being processed
	br.logger.Info("Getting basket for user ID %s", userId)

	// Try to get basket from Redis
	basket, err := br.redisPersistence.Get(userId)
	if err != nil {
		// Try to get basket from MongoDB
		basket, err = br.mongoPersistence.Get(userId)
		if err != nil {
			br.logger.Errorf("Failed to get basket for user ID %s: %v", userId, err)
			return nil, err
		}

		// Cache basket in Redis
		basket, err = br.redisPersistence.Update(basket)
		if err != nil {
			br.logger.Errorf("Failed to update basket in Redis for user ID %s: %v", userId, err)
			return nil, err
		}
	}

	// Log the basket being returned
	br.logger.Info("Returned basket %+v for user ID %s", basket, userId)

	return basket, nil
}

func (br *BasketRepositoryImpl) UpdateBasket(basket *entities.Basket) (*entities.Basket, error) {

	userId := basket.UserID

	// Log the user ID being processed
	br.logger.Info("Updating basket for user ID %s", userId)

	// Update basket in Redis
	update_basket, err := br.redisPersistence.Update(basket)
	if err != nil {
		br.logger.Errorf("Failed to update basket in Redis for user ID %s: %v", userId, err)
		return nil, err
	}

	// Update basket in MongoDB
	if _, err := br.mongoPersistence.Update(basket); err != nil {
		// Rollback Redis basket update on error
		oldCart, err := br.redisPersistence.Get(userId)
		if err != nil {
			br.logger.Errorf("Failed to get basket in Redis for user ID %s: %v", userId, err)
			return nil, err
		}
		if _, err := br.redisPersistence.Update(oldCart); err != nil {
			br.logger.Errorf("Failed to update basket in Redis for user ID %s: %v", userId, err)
			return nil, err
		}

		return nil, err
	}

	// Log the basket being returned
	br.logger.Info("Updated basket %+v for user ID %s", update_basket, userId)

	return update_basket, nil
}

func (br *BasketRepositoryImpl) DeleteBasket(userName string) error {
	// Delete basket from Redis
	if err := br.redisPersistence.Delete(userName); err != nil {
		return err
	}

	// Delete basket from MongoDB
	if err := br.mongoPersistence.Delete(userName); err != nil {
		// // Rollback Redis basket deletion on error
		// oldCart, err := br.redisPersistence.Create(userName)
		// if err != nil {
		// 	return err
		// }
		// if _, err := br.redisPersistence.Update(oldCart); err != nil {
		// 	return err
		// }

		return err
	}

	return nil
}
