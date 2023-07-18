package repositories_test

import (
	"testing"

	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/entities"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/interfaces/persistence"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/repositories"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestCreateBasket(t *testing.T) {
	logger := logrus.New().WithField("module", "BasketRepositoryImpl")
	mockMongoPersistence := &persistence.MockBasketPersistence{}
	mockRedisPersistence := &persistence.MockBasketPersistence{}
	repo := repositories.NewBasketRepositoryImpl(logger, mockMongoPersistence, mockRedisPersistence)

	userID := "user123"
	basket := &entities.Basket{
		UserID: userID,
		// Add other necessary fields
	}

	// Mock the behavior of the mock persistence implementations
	mockMongoPersistence.On("Get", userID).Return(nil, nil)
	mockRedisPersistence.On("Get", userID).Return(nil, nil)
	mockRedisPersistence.On("Create", basket).Return(basket, nil)
	mockMongoPersistence.On("Create", basket).Return(basket, nil)

	createdBasket, err := repo.CreateBasket(basket)

	assert.NoError(t, err)
	assert.Equal(t, basket, createdBasket)

	mockMongoPersistence.AssertExpectations(t)
	mockRedisPersistence.AssertExpectations(t)
}

// Implement other test cases for GetBasket, UpdateBasket, and DeleteBasket
