package repositories

import (
	"testing"

	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/entities"
	"github.com/huavanthong/microservice-golang/src/Services/Basket/internal/domain/repositories"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock for BasketPersistence
type MockBasketPersistence struct {
	mock.Mock
}

func (m *MockBasketPersistence) Get(userID string) (*entities.Basket, error) {
	args := m.Called(userID)
	return args.Get(0).(*entities.Basket), args.Error(1)
}

func (m *MockBasketPersistence) Create(basket *entities.Basket) (*entities.Basket, error) {
	args := m.Called(basket)
	return args.Get(0).(*entities.Basket), args.Error(1)
}

func (m *MockBasketPersistence) Update(basket *entities.Basket) (*entities.Basket, error) {
	args := m.Called(basket)
	return args.Get(0).(*entities.Basket), args.Error(1)
}

func (m *MockBasketPersistence) Delete(userID string) error {
	args := m.Called(userID)
	return args.Error(0)
}

// Test cases
func TestCreateBasket(t *testing.T) {
	logger := logrus.NewEntry(logrus.StandardLogger())

	mongoMock := new(MockBasketPersistence)
	redisMock := new(MockBasketPersistence)

	repo := repositories.NewBasketRepositoryImpl(logger, mongoMock, redisMock)

	basket := &entities.Basket{UserID: "user1"}

	mongoMock.On("Get", "user1").Return((*entities.Basket)(nil), nil)
	redisMock.On("Get", "user1").Return((*entities.Basket)(nil), nil)
	redisMock.On("Create", basket).Return(basket, nil)
	mongoMock.On("Create", basket).Return(basket, nil)

	result, err := repo.CreateBasket(basket)
	assert.NoError(t, err)
	assert.Equal(t, basket, result)
}

func TestGetBasket(t *testing.T) {
	logger := logrus.NewEntry(logrus.StandardLogger())

	mongoMock := new(MockBasketPersistence)
	redisMock := new(MockBasketPersistence)

	repo := repositories.NewBasketRepositoryImpl(logger, mongoMock, redisMock)

	basket := &entities.Basket{UserID: "user1"}
	redisMock.On("Get", "user1").Return(basket, nil)
	mongoMock.On("Get", "user1").Return(basket, nil)
	redisMock.On("Update", basket).Return(basket, nil)

	result, err := repo.GetBasket("user1")
	assert.NoError(t, err)
	assert.Equal(t, basket, result)
}

func TestUpdateBasket(t *testing.T) {
	// Create logger
	logger := logrus.NewEntry(logrus.StandardLogger())

	// Create mocks for MongoDB and Redis persistence
	mongoMock := new(MockBasketPersistence)
	redisMock := new(MockBasketPersistence)

	// Create repository instance with mocked persistence
	repo := repositories.NewBasketRepositoryImpl(logger, mongoMock, redisMock)

	// Prepare test data
	basket := &entities.Basket{UserID: "user1"}

	// Setup expectations for Redis mock
	redisMock.On("Update", basket).Return(basket, nil)
	mongoMock.On("Update", basket).Return(basket, nil)

	// Run the repository function being tested
	result, err := repo.UpdateBasket(basket)

	// Verify behavior
	assert.NoError(t, err)          // Assert no error (as Redis rollback should work)
	assert.Equal(t, basket, result) // Assert the returned basket matches

	// Assert that all expectations were met
	redisMock.AssertExpectations(t)
	mongoMock.AssertExpectations(t)
}

func TestDeleteBasket(t *testing.T) {
	logger := logrus.NewEntry(logrus.StandardLogger())

	mongoMock := new(MockBasketPersistence)
	redisMock := new(MockBasketPersistence)

	repo := repositories.NewBasketRepositoryImpl(logger, mongoMock, redisMock)

	redisMock.On("Delete", "user1").Return(nil)
	mongoMock.On("Delete", "user1").Return(nil)
	redisMock.On("Get", "user1").Return((*entities.Basket)(nil), nil)

	err := repo.DeleteBasket("user1")
	assert.NoError(t, err)
}
