package repositories

import (
	"context"
	"encoding/json"
	"fmt"

	redis "github.com/go-redis/redis/v8"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/entities"
)

type RedisBasketRepository struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisBasketRepository(client *redis.Client, ctx context.Context) *RedisBasketRepository {
	return &RedisBasketRepository{
		client: client,
		ctx:    ctx,
	}
}

func (r *RedisBasketRepository) GetBasket(userName string) (*entities.ShoppingCart, error) {
	val, err := r.client.Get(r.ctx, userName).Result()
	if err != nil {
		return nil, fmt.Errorf("could not get basket for user %s: %v", userName, err)
	}

	basket := &entities.ShoppingCart{}
	if err := json.Unmarshal([]byte(val), basket); err != nil {
		return nil, fmt.Errorf("could not unmarshal basket for user %s: %v", userName, err)
	}

	return basket, nil
}

func (r *RedisBasketRepository) UpdateBasket(basket *entities.ShoppingCart) error {
	basketJSON, err := json.Marshal(basket)
	if err != nil {
		return fmt.Errorf("could not marshal basket: %v", err)
	}

	if err := r.client.Set(r.ctx, basket.UserName, basketJSON, 0).Err(); err != nil {
		return fmt.Errorf("could not set basket for user %s: %v", basket.UserName, err)
	}

	return nil
}

func (r *RedisBasketRepository) DeleteBasket(userName string) error {
	if err := r.client.Del(r.ctx, userName).Err(); err != nil {
		return fmt.Errorf("could not delete basket for user %s: %v", userName, err)
	}

	return nil
}
