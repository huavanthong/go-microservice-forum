package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type HttpHandler struct {
	CommandBus      CommandBus
	ProductQuery    ProductQuery
	OrderQuery      OrderQuery
	DiscountService DiscountService
}

func (h *HttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var cmd CreateOrderCommand
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.CommandBus.Send(context.Background(), cmd); err != nil {
		http.Error(w, "failed to create order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *HttpHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	product, err := h.ProductQuery.GetProduct(id)
	if err != nil {
		http.Error(w, "failed to get product", http.StatusInternalServerError)
		return
	}

	if product == nil {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func (h *HttpHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	order, err := h.OrderQuery.GetOrder(id)
	if err != nil {
		http.Error(w, "failed to get order", http.StatusInternalServerError)
		return
	}

	if order == nil {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(order)
}

func (h *HttpHandler) GetDiscount(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	discount, err := h.DiscountService.GetDiscount(id)
	if err != nil {
		http.Error(w, "failed to get discount", http.StatusInternalServerError)
		return
	}

	if discount == nil {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(discount)
}

func (h *HttpHandler) ApplyDiscount(w http.ResponseWriter, r *http.Request) {
	var cmd ApplyDiscountCommand
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.CommandBus.Send(context.Background(), cmd); err != nil {
		http.Error(w, "failed to apply discount", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
