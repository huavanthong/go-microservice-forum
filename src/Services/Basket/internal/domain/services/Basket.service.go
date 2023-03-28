package services

import (
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/entities"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/infrastructure/persistence/mongodb"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/infrastructure/persistence/redis"
)

type BasketService struct {
	redisRepo *redis.RedisBasketRepository
	mongoRepo *mongodb.MongoDBBasketRepository
}

func NewBasketService(redisRepo *redis.RedisBasketRepository, mongoRepo *mongodb.MongoDBBasketRepository) *BasketService {
	return &BasketService{redisRepo, mongoRepo}
}

func (s *BasketService) CreateBasket(userName string) (*entities.ShoppingCart, error) {
	// Create basket in Redis
	basket, err := s.redisRepo.Create(userName)
	if err != nil {
		return nil, err
	}

	// Create basket in MongoDB
	_, err = s.mongoRepo.Create(userName)
	if err != nil {
		// Rollback Redis basket creation on error
		s.redisRepo.Delete(userName)
		return nil, err
	}

	return basket, nil
}

func (s *BasketService) GetBasket(userName string) (*entities.ShoppingCart, error) {
	// Try to get basket from Redis
	basket, err := s.redisRepo.GetByUserName(userName)
	if err != nil {
		// Try to get basket from MongoDB
		basket, err = s.mongoRepo.GetByUserName(userName)
		if err != nil {
			return nil, err
		}

		// Cache basket in Redis
		basket, err = s.redisRepo.Update(basket)
		if err != nil {
			return nil, err
		}
	}

	return basket, nil
}

func (s *BasketService) UpdateBasket(userName string, cart *entities.ShoppingCart) (*entities.ShoppingCart, error) {
	// Update basket in Redis
	if _, err := s.redisRepo.Update(cart); err != nil {
		return nil, err
	}

	// Update basket in MongoDB
	if _, err := s.mongoRepo.Update(cart); err != nil {
		// Rollback Redis basket update on error
		oldCart, err := s.redisRepo.GetByUserName(userName)
		if err != nil {
			return nil, err
		}
		if _, err := s.redisRepo.Update(oldCart); err != nil {
			return nil, err
		}

		return nil, err
	}

	return cart, nil
}

func (s *BasketService) DeleteBasket(userName string) error {
	// Delete basket from Redis
	if err := s.redisRepo.Delete(userName); err != nil {
		return err
	}

	// Delete basket from MongoDB
	if err := s.mongoRepo.Delete(userName); err != nil {
		// Rollback Redis basket deletion on error
		oldCart, err := s.redisRepo.Create(userName)
		if err != nil {
			return err
		}
		if _, err := s.redisRepo.Update(oldCart); err != nil {
			return err
		}

		return err
	}

	return nil
}
