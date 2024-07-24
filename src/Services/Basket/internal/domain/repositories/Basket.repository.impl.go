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

	// Extract user Id from request
	userId := basket.UserID

	// Log the user ID being processed
	br.logger.Infof("Creating basket for user ID: %s", userId)

	// Try to get basket from MongoDB
	find_basket, err := br.mongoPersistence.Get(userId)
	create_basket := find_basket
	if err != nil {
		br.logger.Errorf("Failed to get basket from MongoDB: %s", err.Error())

		// Try to get basket from Redis
		if find_basket, err = br.redisPersistence.Get(userId); err != nil {
			br.logger.Errorf("Failed to get basket from Redis: %s", err.Error())
			return nil, err
		}

		create_basket = find_basket
		if find_basket == nil {
			// Basket not found in MongoDB and Redis, create a new one
			if create_basket, err = br.redisPersistence.Create(basket); err != nil {
				br.logger.Errorf("Failed to create basket in Redis: %s", err.Error())
				return nil, err
			}

			// Create basket in MongoDB
			if create_basket, err = br.mongoPersistence.Create(basket); err != nil {
				// Rollback Redis basket creation on error
				br.redisPersistence.Delete(userId)
				br.logger.Errorf("Failed to create basket in MongoDB: %s", err.Error())
				return nil, err
			}
		}
	}

	// Log the basket being returned
	br.logger.Infof("Created basket %+v for user ID %s", create_basket, create_basket.UserID)

	return create_basket, nil
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

	get_basket := find_basket
	if get_basket == nil {

		// Try to get basket from MongoDB
		get_basket, err = br.mongoPersistence.Get(userId)
		if err != nil {
			br.logger.Errorf("Failed to get basket for user ID %s: %v", userId, err)
			return nil, err
		}

		// Cache basket in Redis
		get_basket, err = br.redisPersistence.Update(get_basket)
		if err != nil {
			br.logger.Errorf("Failed to update basket in Redis for user ID %s: %v", userId, err)
			return nil, err
		}
	}

	// Log the basket being returned
	br.logger.Infof("Returned basket %+v for user ID %s ", get_basket, userId)

	return get_basket, nil
}

func (br *BasketRepositoryImpl) UpdateBasket(basket *entities.Basket) (*entities.Basket, error) {

	userId := basket.UserID

	// Log the user ID being processed
	br.logger.Infof("Updating basket for user ID %s: ", userId)

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
	br.logger.Infof("Updated basket %+v for user ID %s", update_basket, userId)

	return update_basket, nil
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
		get_basket, err := br.redisPersistence.Get(userId)
		if err != nil {
			br.logger.Errorf("Failed to get basket in Redis for user ID %s during rollback: %v", userId, err)
			return err
		}
		if get_basket == nil {
			br.logger.Warnf("No basket found in Redis for user ID %s during rollback", userId)
			return err
		}
		if _, err := br.redisPersistence.Create(get_basket); err != nil {
			br.logger.Errorf("Failed to create basket in Redis for user ID %s during rollback: %v", userId, err)
			return err
		}
		if _, err := br.redisPersistence.Update(get_basket); err != nil {
			br.logger.Errorf("Failed to update basket in Redis for user ID %s during rollback: %v", userId, err)
			return err
		}

		br.logger.Infof("Rollback succeeded for user ID %s", userId)
		return err
	}

	// Log the basket being returned
	br.logger.Infof("Deleted basket success for user ID %s", userId)

	return nil
}
