package httpserver

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/your_username/order-service/internal/order"
)

type HttpHandler struct {
	commandBus CommandBus
	query      order.Query
}

func NewHttpHandler(commandBus CommandBus, query order.Query) *HttpHandler {
	return &HttpHandler{
		commandBus: commandBus,
		query:      query,
	}
}

func (h *HttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req order.CreateOrderRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	cmd := order.CreateOrderCommand{
		OrderID:      req.OrderID,
		CustomerID:   req.CustomerID,
		ProductID:    req.ProductID,
		Quantity:     req.Quantity,
		Price:        req.Price,
		TotalAmount:  req.TotalAmount,
		ShippingCost: req.ShippingCost,
	}

	err = h.commandBus.Publish(r.Context(), cmd)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *HttpHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	orderID := mux.Vars(r)["orderID"]

	cmd := order.DeleteOrderCommand{
		OrderID: orderID,
	}

	err := h.commandBus.Publish(r.Context(), cmd)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *HttpHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	orderID := mux.Vars(r)["orderID"]

	var req order.UpdateOrderRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	cmd := order.UpdateOrderCommand{
		OrderID:      orderID,
		ProductID:    req.ProductID,
		Quantity:     req.Quantity,
		Price:        req.Price,
		TotalAmount:  req.TotalAmount,
		ShippingCost: req.ShippingCost,
	}

	err = h.commandBus.Publish(r.Context(), cmd)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
