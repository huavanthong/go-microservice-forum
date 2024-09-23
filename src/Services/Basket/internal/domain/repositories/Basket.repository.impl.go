package repositories

import (
	"fmt"

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

	// Extract user Id from request
	userId := basket.UserID

	// Log the user ID being processed
	br.logger.Infof("Creating basket for user ID: %s", userId)

	// Try to get basket from MongoDB
	find_basket, err := br.mongoPersistence.Get(userId)
	if err != nil {
		br.logger.Errorf("Failed to get basket from MongoDB: %s", err.Error())
		return nil, err
	}

	if find_basket == nil {
		// Try to get basket from Redis
		if find_basket, err = br.redisPersistence.Get(userId); err != nil {
			br.logger.Errorf("Failed to get basket from Redis: %s", err.Error())
			return nil, err
		}
	}

	if find_basket == nil {
		// Basket not found in both MongoDB and Redis, create a new one
		create_basket, err := br.redisPersistence.Create(basket)
		if err != nil {
			br.logger.Errorf("Failed to create basket in Redis: %s", err.Error())
			return nil, err
		}
		// Log the basket being returned
		br.logger.Infof("Created basket in Redis %+v for user ID %s", create_basket, create_basket.UserID)

		create_basket, err = br.mongoPersistence.Create(basket)
		if err != nil {
			br.redisPersistence.Delete(userId)
			br.logger.Errorf("Failed to create basket in MongoDB: %s", err.Error())
			return nil, err
		}

		// Log the basket being returned
		br.logger.Infof("Created basket in MongoDB %+v for user ID %s", create_basket, create_basket.UserID)

		return create_basket, nil
	}

	// Log the basket being returned
	br.logger.Infof("Get basket %+v from exist User ID  %s", find_basket, find_basket.UserID)

	return find_basket, nil
}

func (br *BasketRepositoryImpl) GetBasket(userId string) (*entities.Basket, error) {

	// Log the user ID being processed
	br.logger.Infof("Getting basket for user ID: %s", userId)

	// Try to get basket from Redis
	find_basket, err := br.redisPersistence.Get(userId)
	if err != nil {
		br.logger.Errorf("Failed to get basket from Redis: %s", err.Error())
		return nil, err
	}

	// Log the basket being returned
	br.logger.Infof("Returned basket in Redis %+v for user ID %s ", find_basket, userId)

	if find_basket == nil {

		// Try to get basket from MongoDB
		find_basket, err = br.mongoPersistence.Get(userId)
		if err != nil {
			br.logger.Errorf("Failed to get basket for user ID %s: %v", userId, err)
			return nil, err
		}

		// Log the basket being returned
		br.logger.Infof("Returned basket in MongoDB %+v for user ID %s ", find_basket, userId)

		// Cache basket in Redis
		cache_basket, err := br.redisPersistence.Update(find_basket)
		if err != nil {
			br.logger.Errorf("Failed to update basket in Redis for user ID %s: %v", userId, err)
			return nil, err
		}

		// Log the basket being returned
		br.logger.Infof("Caching basket in Redis %+v for user ID %s ", cache_basket, userId)
	}

	return find_basket, nil
}

func (br *BasketRepositoryImpl) UpdateBasket(basket *entities.Basket) (*entities.Basket, error) {

	userId := basket.UserID

	// Log the user ID being processed
	br.logger.Infof("Updating basket for user ID %s: ", userId)

	// Update basket in Redis
	update_redis_basket, err := br.redisPersistence.Update(basket)
	if err != nil {
		br.logger.Errorf("Failed to update basket in Redis for user ID %s: %v", userId, err)
		return nil, err
	}
	// Log the basket being returned
	br.logger.Infof("Updated basket in Redis %+v for user ID %s", update_redis_basket, userId)

	// Update basket in MongoDB
	updateMongoBasket, err := br.mongoPersistence.Update(basket)
	if err != nil {
		br.logger.Errorf("Failed to update basket in MongoDB for user ID %s: %v", userId, err)

		// Rollback Redis basket update on error
		oldCart, getErr := br.redisPersistence.Get(userId)
		if getErr != nil {
			br.logger.Errorf("Failed to get basket in Redis for user ID %s during rollback: %v", userId, getErr)
			return nil, getErr
		}

		if _, rollbackErr := br.redisPersistence.Update(oldCart); rollbackErr != nil {
			br.logger.Errorf("Failed to rollback basket in Redis for user ID %s: %v", userId, rollbackErr)
			return nil, rollbackErr
		}

		return nil, err
	}
	br.logger.Infof("Updated basket in MongoDB %+v for user ID %s", updateMongoBasket, userId)

	// Return the updated basket
	return updateMongoBasket, nil
}

func (br *BasketRepositoryImpl) DeleteBasket(userId string) error {

	// Log the user ID being processed
	br.logger.Infof("Deleting basket for user ID %s: ", userId)

	// Delete basket from Redis
	if err := br.redisPersistence.Delete(userId); err != nil {
		br.logger.Errorf("Failed to delete basket in Redis for user ID %s: %v", userId, err)
		return err
	}

	// Delete basket from MongoDB
	if err := br.mongoPersistence.Delete(userId); err != nil {
		// Rollback Redis basket deletion on error
		br.logger.Warnf("Failed to delete basket in MongoDB for user ID %s, rolling back Redis deletion", userId)

		// Attempt to get the basket from Redis for rollback
		get_basket, err := br.redisPersistence.Get(userId)
		if err != nil {
			br.logger.Errorf("Failed to get basket in Redis for user ID %s during rollback: %v", userId, err)
			return fmt.Errorf("MongoDB delete failed, rollback failed: %v", err)
		}
		if get_basket == nil {
			br.logger.Warnf("No basket found in Redis for user ID %s during rollback", userId)
			return fmt.Errorf("MongoDB delete failed, no basket in Redis to roll back")
		}

		// Restore the basket in Redis
		if _, err := br.redisPersistence.Update(get_basket); err != nil {
			br.logger.Errorf("Failed to update basket in Redis for user ID %s during rollback: %v", userId, err)
			return fmt.Errorf("MongoDB delete failed, Redis rollback failed: %v", err)
		}

		br.logger.Infof("Rollback succeeded for user ID %s", userId)
		return fmt.Errorf("MongoDB delete failed, rollback succeeded")
	}

	// Log the basket being returned
	br.logger.Infof("Deleted basket success for user ID %s", userId)

	return nil
}
