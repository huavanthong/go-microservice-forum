package command

import (
	"context"
	"fmt"
	"log"
)

type CheckoutOrderCommandHandler struct {
	orderRepository IOrderRepository
	mapper          IMapper
	emailService    IEmailService
	logger          *log.Logger
}

func NewCheckoutOrderCommandHandler(orderRepository IOrderRepository, mapper IMapper, emailService IEmailService, logger *log.Logger) *CheckoutOrderCommandHandler {
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

func (h *CheckoutOrderCommandHandler) Handle(ctx context.Context, req *CheckoutOrderCommand) (int, error) {
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
