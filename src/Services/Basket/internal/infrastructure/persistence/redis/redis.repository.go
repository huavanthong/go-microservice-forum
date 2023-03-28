package redis

import (
	"context"
	"encoding/json"
	"fmt"

	redis "github.com/go-redis/redis/v8"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/entities"
)

// Define struct for RedisBasketRepository
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
	key := fmt.Sprintf("basket:%s", userName)
	basket := &entities.ShoppingCart{UserName: userName}

	data, err := json.Marshal(basket)
	if err != nil {
		return nil, fmt.Errorf("failed to create basket for user %s: %v", userName, err)
	}

	err = r.client.Set(r.ctx, key, data, 0).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to create basket for user %s: %v", userName, err)
	}

	return basket, nil
}

func (r *RedisBasketRepository) GetByUserName(userName string) (*entities.ShoppingCart, error) {
	key := fmt.Sprintf("basket:%s", userName)

	data, err := r.client.Get(r.ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get basket for user %s: %v", userName, err)
	}

	basket := &entities.ShoppingCart{}
	err = json.Unmarshal(data, basket)
	if err != nil {
		return nil, fmt.Errorf("failed to get basket for user %s: %v", userName, err)
	}

	return basket, nil
}

func (r *RedisBasketRepository) Update(basket *entities.ShoppingCart) (*entities.ShoppingCart, error) {
	key := fmt.Sprintf("basket:%s", basket.UserName)

	data, err := json.Marshal(basket)
	if err != nil {
		return nil, fmt.Errorf("failed to update basket for user %s: %v", basket.UserName, err)
	}

	err = r.client.Set(r.ctx, key, data, 0).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to update basket for user %s: %v", basket.UserName, err)
	}

	return basket, nil
}

func (r *RedisBasketRepository) Delete(userName string) error {
	key := fmt.Sprintf("basket:%s", userName)

	err := r.client.Del(r.ctx, key).Err()
	if err != nil {
		return fmt.Errorf("failed to delete basket for user %s: %v", userName, err)
	}

	return nil
}
