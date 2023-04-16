package handler

import (
	"context"
	"fmt"
	"log"

	"github.com/huavanthong/microservice-golang/src/Services/Ordering/internal/domain/repository"
	"github.com/huavanthong/microservice-golang/src/Services/Ordering/internal/domain/service"
)

type CheckoutOrderCommandHandler struct {
	orderRepository repository.OrderRepository
	mapper          IMapper
	emailService    service.EmailService
	logger          *log.Logger
}

type CheckoutOrderCommandRequest struct {
	UserName      string  `json:"userName"`
	TotalPrice    float64 `json:"totalPrice"`
	FirstName     string  `json:"firstName"`
	LastName      string  `json:"lastName"`
	EmailAddress  string  `json:"emailAddress"`
	AddressLine   string  `json:"addressLine"`
	Country       string  `json:"country"`
	State         string  `json:"state"`
	ZipCode       string  `json:"zipCode"`
	CardName      string  `json:"cardName"`
	CardNumber    string  `json:"cardNumber"`
	Expiration    string  `json:"expiration"`
	CVV           string  `json:"cvv"`
	PaymentMethod int     `json:"paymentMethod"`
}

type CheckoutOrderCommandResponse struct {
	ID string
}

func NewCheckoutOrderCommandHandler(orderRepository repository.OrderRepository, mapper IMapper, emailService IEmailService, logger *log.Logger) *CheckoutOrderCommandHandler {
	if orderRepository == nil {
		panic("orderRepository is nil")
	}
	if mapper == nil {
		panic("mapper is nil")
	}
	if emailService == nil {
		panic("emailService is nil")
	}
	if logger == nil {
		panic("logger is nil")
	}

	return &CheckoutOrderCommandHandler{
		orderRepository: orderRepository,
		mapper:          mapper,
		emailService:    emailService,
		logger:          logger,
	}
}

func (h *CheckoutOrderCommandHandler) Handle(ctx context.Context, req *CheckoutOrderCommandRequest) (int, error) {

	orderEntity := h.mapper.Map(req, new(Order))

	newOrder, err := h.orderRepository.Add(orderEntity)
	if err != nil {
		return 0, fmt.Errorf("failed to add new order: %v", err)
	}

	h.logger.Printf("Order %d is successfully created.", newOrder.Id)

	err = h.sendMail(newOrder)
	if err != nil {
		h.logger.Printf("Order %d failed due to an error with the mail service: %v", newOrder.Id, err)
	}

	return newOrder.Id, nil
}

func (h *CheckoutOrderCommandHandler) sendMail(order *Order) error {
	email := Email{To: "ezozkme@gmail.com", Body: "Order was created.", Subject: "Order was created"}

	err := h.emailService.SendEmail(email)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
