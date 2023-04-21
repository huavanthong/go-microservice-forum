package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"

	redis "github.com/go-redis/redis/v8"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/entities"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/interfaces/persistence"
)

// Define struct for Redis Basket Repository
type RedisBasketPersistence struct {
	logger *logrus.Entry // add logger field
	client *redis.Client
	ctx    context.Context
}

// NewRedisBasketPersistence returns a new instance of RedisBasketPersistence
func NewRedisBasketPersistence(logger *logrus.Entry, client *redis.Client, ctx context.Context) persistence.BasketPersistence {
	return &RedisBasketPersistence{
		logger: logrus.WithField("module", "RedisBasketPersistence"), // set logger field
		client: client,
		ctx:    ctx,
	}
}

func (rbp *RedisBasketPersistence) Create(basket *entities.Basket) (*entities.Basket, error) {

	rbp.logger.Info("Creating a basket for user id %s in Redis", basket.UserID)
	// Concatenate the userId parameter with the string "basket:".
	// This is the key that will be used to store the basket in Redis.
	var userId string = basket.UserID

	key := fmt.Sprintf("basket:%s", userId)

	// Create shopping cart based on user name

	// Serialize the ShoppingCart object into a JSON string
	data, err := json.Marshal(basket)
	if err != nil {
		rbp.logger.Errorf("failed to create basket for user %s: %v", userId, err) // log error message
		return nil, fmt.Errorf("failed to create basket for user %s: %v", userId, err)
	}

	// Set the value for the key in Redis.
	err = rbp.client.Set(rbp.ctx, key, data, 0).Err()
	if err != nil {
		rbp.logger.Errorf("failed to create basket for user %s: %v", userId, err) // log error message
		return nil, fmt.Errorf("failed to create basket for user %s: %v", userId, err)
	}

	rbp.logger.Info("Basket for user %s has been successfully retrieved from Redis", basket.UserID)

	return basket, nil
}

func (rbp *RedisBasketPersistence) Get(userId string) (*entities.Basket, error) {

	rbp.logger.Info("Getting basket by user id %s in Redis ", userId)

	// Generating key
	key := fmt.Sprintf("basket:%s", userId)

	// Retrieves the data associated with the key from Redis
	data, err := rbp.client.Get(rbp.ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			rbp.logger.Warn("Basket not found in Redis")
			return nil, nil
		}

		rbp.logger.Errorf("Failed to get basket from Redis")
		return nil, fmt.Errorf("failed to get basket for user %s: %v", userId, err)
	}

	basket := &entities.Basket{}

	// If data is retrieved successfully,it unmarshals the data into a new instance.
	err = json.Unmarshal(data, basket)
	if err != nil {
		rbp.logger.Errorf("Failed to unmarshal basket data from Redis")
		return nil, fmt.Errorf("failed to get basket for user %s: %v", userId, err)
	}

	rbp.logger.Info("Got a basket by user id %s successfully in Redis", basket.UserID)

	return basket, nil
}

func (rbp *RedisBasketPersistence) Update(basket *entities.Basket) (*entities.Basket, error) {

	rbp.logger.Info("Updating the basket by user id %s in Redis", basket.UserID)

	// Generating key
	key := fmt.Sprintf("basket:%s", basket.UserID)

	// Serialize the ShoppingCart object into a JSON string
	data, err := json.Marshal(basket)
	if err != nil {
		rbp.logger.Errorf("failed to update basket for user %s: %v", basket.UserID, err)
		return nil, fmt.Errorf("failed to update basket for user %s: %v", basket.UserID, err)
	}

	// Set the value for the key in Redis.
	err = rbp.client.Set(rbp.ctx, key, data, 0).Err()
	if err != nil {
		rbp.logger.Errorf("failed to update basket for user %s: %v", basket.UserID, err)
		return nil, fmt.Errorf("failed to update basket for user %s: %v", basket.UserID, err)
	}

	rbp.logger.Info("Updated a basket by user id %s successfully in Redis", basket.UserID)

	return basket, nil
}

func (rbp *RedisBasketPersistence) Delete(userId string) error {

	rbp.logger.Info("Deleting the basket by user id %s in Redis", userId)

	// Generating key
	key := fmt.Sprintf("basket:%s", userId)

	// Delete the data associated with the key from Redis
	err := rbp.client.Del(rbp.ctx, key).Err()
	if err != nil {
		rbp.logger.Errorf("failed to delete basket for user %s: %v", userId, err)
		return fmt.Errorf("failed to delete basket for user %s: %v", userId, err)
	}

	rbp.logger.Info("Deleted a basket by user id %s successfully in Redis", userId)

	return nil
}
