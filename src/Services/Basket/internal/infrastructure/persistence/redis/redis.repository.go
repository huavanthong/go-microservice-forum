package redis

import (
	"context"
	"encoding/json"
	"fmt"

	redis "github.com/go-redis/redis/v8"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/entities"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/interfaces/persistence"
)

// Define struct for Redis Basket Repository
type RedisBasketPersistence struct {
	client *redis.Client
	ctx    context.Context
}

// NewRedisBasketPersistence returns a new instance of RedisBasketPersistence
func NewRedisBasketPersistence(client *redis.Client, ctx context.Context) persistence.BasketPersistence {
	return &RedisBasketPersistence{
		client: client,
		ctx:    ctx,
	}
}

func (rbp *RedisBasketPersistence) Create(userId string) (*entities.Basket, error) {

	if _, err := rbp.client.Ping(rbp.ctx).Result(); err != nil {
		panic(err)
	}
	// Concatenate the userId parameter with the string "basket:".
	// This is the key that will be used to store the basket in Redis.

	key := fmt.Sprintf("basket:%s", userId)

	// Create shopping cart based on user name
	basket := &entities.Basket{UserID: userId}

	// Serialize the ShoppingCart object into a JSON string
	data, err := json.Marshal(basket)
	if err != nil {
		return nil, fmt.Errorf("failed to create basket for user %s: %v", userId, err)
	}

	// Set the value for the key in Redis.
	err = rbp.client.Set(rbp.ctx, key, data, 0).Err()
	fmt.Println("Check redis 3: ", err)
	if err != nil {
		return nil, fmt.Errorf("failed to create basket for user %s: %v", userId, err)
	}

	return basket, nil
}

func (rbp *RedisBasketPersistence) Get(userId string) (*entities.Basket, error) {

	// Generating key
	key := fmt.Sprintf("basket:%s", userId)

	// Retrieves the data associated with the key from Redis
	data, err := rbp.client.Get(rbp.ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get basket for user %s: %v", userId, err)
	}

	basket := &entities.Basket{}

	// If data is retrieved successfully,it unmarshals the data into a new instance.
	err = json.Unmarshal(data, basket)
	if err != nil {
		return nil, fmt.Errorf("failed to get basket for user %s: %v", userId, err)
	}

	return basket, nil
}

func (rbp *RedisBasketPersistence) Update(basket *entities.Basket) (*entities.Basket, error) {
	// Generating key
	key := fmt.Sprintf("basket:%s", basket.UserName)

	// Serialize the ShoppingCart object into a JSON string
	data, err := json.Marshal(basket)
	if err != nil {
		return nil, fmt.Errorf("failed to update basket for user %s: %v", basket.UserName, err)
	}

	// Set the value for the key in Redis.
	err = rbp.client.Set(rbp.ctx, key, data, 0).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to update basket for user %s: %v", basket.UserName, err)
	}

	return basket, nil
}

func (rbp *RedisBasketPersistence) Delete(userName string) error {
	// Generating key
	key := fmt.Sprintf("basket:%s", userName)

	// Delete the data associated with the key from Redis
	err := rbp.client.Del(rbp.ctx, key).Err()
	if err != nil {
		return fmt.Errorf("failed to delete basket for user %s: %v", userName, err)
	}

	return nil
}
