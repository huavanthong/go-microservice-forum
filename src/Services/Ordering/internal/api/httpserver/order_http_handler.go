package httpserver

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/huavanthong/microservice-golang/src/Services/Ordering/internal/application/commands"
	"github.com/huavanthong/microservice-golang/src/Services/Ordering/internal/application/quries"
)

type HttpHandler struct {
	commandBus commands.CommandBus
	query      quries.Query
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

	err = h.commandBus.Dispatch(r.Context(), cmd)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *HttpHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	var req order.GetOrderRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	cmd := order.GetOrderCommand{
		OrderID:      req.OrderID,
		CustomerID:   req.CustomerID,
		ProductID:    req.ProductID,
		Quantity:     req.Quantity,
		Price:        req.Price,
		TotalAmount:  req.TotalAmount,
		ShippingCost: req.ShippingCost,
	}

	err = h.commandBus.Dispatch(r.Context(), cmd)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *HttpHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	orderID := mux.Vars(r)["orderID"]

	cmd := order.DeleteOrderCommand{
		OrderID: orderID,
	}

	err := h.commandBus.Dispatch(r.Context(), cmd)
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

	err = h.commandBus.Dispatch(r.Context(), cmd)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
