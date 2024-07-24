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

	mongoMock.On("Get", "user1").Return((*entities.Basket)(nil), assert.AnError)
	redisMock.On("Get", "user1").Return((*entities.Basket)(nil), assert.AnError)
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
	logger := logrus.NewEntry(logrus.StandardLogger())

	mongoMock := new(MockBasketPersistence)
	redisMock := new(MockBasketPersistence)

	repo := repositories.NewBasketRepositoryImpl(logger, mongoMock, redisMock)

	basket := &entities.Basket{UserID: "user1"}
	redisMock.On("Update", basket).Return(basket, nil)
	mongoMock.On("Update", basket).Return(nil, assert.AnError)
	redisMock.On("Get", "user1").Return(basket, nil)
	redisMock.On("Create", basket).Return(basket, nil)
	redisMock.On("Update", basket).Return(basket, nil)

	result, err := repo.UpdateBasket(basket)
	assert.NoError(t, err)
	assert.Equal(t, basket, result)
}

func TestDeleteBasket(t *testing.T) {
	logger := logrus.NewEntry(logrus.StandardLogger())

	mongoMock := new(MockBasketPersistence)
	redisMock := new(MockBasketPersistence)

	repo := repositories.NewBasketRepositoryImpl(logger, mongoMock, redisMock)

	redisMock.On("Delete", "user1").Return(nil)
	mongoMock.On("Delete", "user1").Return(assert.AnError)
	redisMock.On("Get", "user1").Return(&entities.Basket{UserID: "user1"}, nil)
	redisMock.On("Create", &entities.Basket{UserID: "user1"}).Return(&entities.Basket{UserID: "user1"}, nil)
	redisMock.On("Update", &entities.Basket{UserID: "user1"}).Return(&entities.Basket{UserID: "user1"}, nil)

	err := repo.DeleteBasket("user1")
	assert.Error(t, err)
}
