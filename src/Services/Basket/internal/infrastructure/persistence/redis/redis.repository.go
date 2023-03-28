package redis

import (
	"context"
	"encoding/json"
	"fmt"

	redis "github.com/go-redis/redis/v8"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/entities"
)

// Define struct for Redis Basket Repository
type RedisBasketRepository struct {
	client *redis.Client
	ctx    context.Context
}

// NewRedisBasketRepository returns a new instance of RedisBasketRepository
func NewRedisBasketRepository(client *redis.Client, ctx context.Context) *RedisBasketRepository {
	return &RedisBasketRepository{
		client: client,
		ctx:    ctx,
	}
}

func (r *RedisBasketRepository) Create(userName string) (*entities.ShoppingCart, error) {

	// Concatenate the userName parameter with the string "basket:".
	// This is the key that will be used to store the basket in Redis.
	key := fmt.Sprintf("basket:%s", userName)

	// Create shopping cart based on user name
	basket := &entities.ShoppingCart{UserName: userName}

	// Serialize the ShoppingCart object into a JSON string
	data, err := json.Marshal(basket)
	if err != nil {
		return nil, fmt.Errorf("failed to create basket for user %s: %v", userName, err)
	}

	// Set the value for the key in Redis.
	err = r.client.Set(r.ctx, key, data, 0).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to create basket for user %s: %v", userName, err)
	}

	return basket, nil
}

func (r *RedisBasketRepository) GetByUserName(userName string) (*entities.ShoppingCart, error) {

	// Generating key
	key := fmt.Sprintf("basket:%s", userName)

	// Retrieves the data associated with the key from Redis
	data, err := r.client.Get(r.ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get basket for user %s: %v", userName, err)
	}

	basket := &entities.ShoppingCart{}

	// If data is retrieved successfully,it unmarshals the data into a new instance.
	err = json.Unmarshal(data, basket)
	if err != nil {
		return nil, fmt.Errorf("failed to get basket for user %s: %v", userName, err)
	}

	return basket, nil
}

func (r *RedisBasketRepository) Update(basket *entities.ShoppingCart) (*entities.ShoppingCart, error) {
	// Generating key
	key := fmt.Sprintf("basket:%s", basket.UserName)

	// Serialize the ShoppingCart object into a JSON string
	data, err := json.Marshal(basket)
	if err != nil {
		return nil, fmt.Errorf("failed to update basket for user %s: %v", basket.UserName, err)
	}

	// Set the value for the key in Redis.
	err = r.client.Set(r.ctx, key, data, 0).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to update basket for user %s: %v", basket.UserName, err)
	}

	return basket, nil
}

func (r *RedisBasketRepository) Delete(userName string) error {
	// Generating key
	key := fmt.Sprintf("basket:%s", userName)

	// Delete the data associated with the key from Redis
	err := r.client.Del(r.ctx, key).Err()
	if err != nil {
		return fmt.Errorf("failed to delete basket for user %s: %v", userName, err)
	}

	return nil
}
