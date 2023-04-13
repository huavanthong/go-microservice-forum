package constracts_infrastructure

import (
	"github.com/huavanthong/microservice-golang/src/Services/Ordering/internal/application/models"
)

type IEmailService interface {
	SendEmail(email models.Email)
}
