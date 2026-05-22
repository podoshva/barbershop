package handlers

import (
	"encoding/json"
	"main/internal/dto"
	"main/internal/handler/httputils"
	"net/http"
)

func (h *Handlers) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var dto dto.CreateOrder
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := h.OrderService.Create(r.Context(), dto); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	httputils.WriteJSON(w, 201, dto)
}

func (h *Handlers) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	id, err := httputils.GetID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := h.BranchService.Delete(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	httputils.WriteJSON(w, 204, nil)
}

func (h *Handlers) GetOrder(w http.ResponseWriter, r *http.Request) {
	id, err := httputils.GetID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	order, err := h.OrderService.Get(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	httputils.WriteJSON(w, 200, order)
}

func (h *Handlers) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.OrderService.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	httputils.WriteJSON(w, 200, orders)
}

func (h *Handlers) SetOrderStatus(w http.ResponseWriter, r *http.Request) {
	id, err := httputils.GetID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var dto dto.SetOrderStatus
	dto.ID = id
	if err = json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	order, err := h.OrderService.SetStatus(r.Context(), dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	httputils.WriteJSON(w, 200, order)
}
