package handler

import (
	"log"
	"net/http"

	"github.com/huavanthong/microservice-golang/src/Services/Ordering/internal/domain/repository"
)

type DeleteOrderCommandHandler struct {
	OrderRepository repository.OrderRepository
	Logger          *log.Logger
}

func NewDeleteOrderCommandHandler(orderRepository repository.OrderRepository, logger *log.Logger) *DeleteOrderCommandHandler {
	return &DeleteOrderCommandHandler{
		OrderRepository: orderRepository,
		Logger:          logger,
	}
}

func (h *DeleteOrderCommandHandler) Handle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["id"]

	orderToDelete, err := h.OrderRepository.GetById(orderID)
	if err != nil {
		if err == ErrOrderNotFound {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		h.Logger.Printf("Error while getting order: %v\n", err)
		return
	}

	err = h.OrderRepository.Delete(orderToDelete)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.Logger.Printf("Error while deleting order: %v\n", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	h.Logger.Printf("Order %s is successfully deleted.", orderToDelete.ID)
}
