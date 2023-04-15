package command

import (
	"errors"
	"log"
	"reflect"
)

type UpdateOrderCommandHandler struct {
	orderRepository IOrderRepository
	mapper          IMapper
	logger          *log.Logger
}

func NewUpdateOrderCommandHandler(orderRepository IOrderRepository, mapper IMapper, logger *log.Logger) *UpdateOrderCommandHandler {
	return &UpdateOrderCommandHandler{orderRepository: orderRepository, mapper: mapper, logger: logger}
}

func (h *UpdateOrderCommandHandler) Handle(request *UpdateOrderCommand) error {
	orderToUpdate, err := h.orderRepository.GetByIdAsync(request.Id)
	if err != nil {
		return err
	}

	if orderToUpdate == nil {
		return errors.New("Order not found")
	}

	err = h.mapper.Map(request, orderToUpdate, reflect.TypeOf(request), reflect.TypeOf(orderToUpdate))
	if err != nil {
		return err
	}

	err = h.orderRepository.UpdateAsync(orderToUpdate)
	if err != nil {
		return err
	}

	h.logger.Println("Order", orderToUpdate.Id, "is successfully updated.")

	return nil
}
