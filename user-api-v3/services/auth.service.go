package services

import "github.com/huavanthong/microservice-golang/user-api-v3/models"

type AuthService interface {
	SignUpUser(*models.SignUpInput) (*models.DBResponse, error)
	SignInUser(*models.SignInInput) (*models.DBResponse, error)
}
