package repositories

import (
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/entities"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/infrastructure/persistence/mongodb"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/infrastructure/redis"
)

type BasketRepository struct {
	redisRepo redis.RedisBasketRepository
	mongoRepo mongodb.MongoBasketRepository
}

func NewBasketRepository(redisRepo redis.RedisBasketRepository) *BasketRepository {
	return &BasketRepository{redisRepo: redisRepo}
}

func (repo *BasketRepository) GetBasket(userName string) (*entities.ShoppingCart, error) {
	return repo.redisRepo.GetBasket(userName)
}

func (repo *BasketRepository) UpdateBasket(basket *entities.ShoppingCart) (*entities.ShoppingCart, error) {
	return repo.redisRepo.UpdateBasket(basket)
}

func (repo *BasketRepository) DeleteBasket(userName string) error {
	return repo.redisRepo.DeleteBasket(userName)
}
