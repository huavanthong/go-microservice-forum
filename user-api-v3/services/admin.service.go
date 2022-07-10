package services

import "github.com/huavanthong/microservice-golang/user-api-v3/models"

type AuthService interface {
	GetAllUsers(*models.SignUpInput) (*models.DBResponse, error)
	GetUserByID(*models.SignInInput) (*models.DBResponse, error)
	GetUserByParam(*models.SignInInput) (*models.DBResponse, error)
}
