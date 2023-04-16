package services

import (
	"github.com/huavanthong/microservice-golang/src/Services/Ordering/internal/application/models"
)

type EmailService interface {
	SendEmail(email models.Email)
}
