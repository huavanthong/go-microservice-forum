package services

import "github.com/huavanthong/microservice-golang/user-api-v3/models"

type AdminService interface {
	GetAllUsers(page int, limit int) ([]*models.User, error)
	GetUserByID(id string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}
