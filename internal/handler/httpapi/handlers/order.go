package handlers

import (
	"encoding/json"
	"main/internal/dto"
	"main/internal/handler/httputils"
	"net/http"
	"time"
)

func (h *Handlers) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var body struct {
		ProfileID     int64     `json:"profile_id,omitempty"`
		BranchID      int64     `json:"branch_id,omitempty"`
		Date          time.Time `json:"date,omitempty"`
		CustomerPhone string    `json:"customer_phone,omitempty"`
		Description   string    `json:"description,omitempty"`
		Status        string    `json:"status,omitempty"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := h.OrderService.Create(r.Context(), dto.CreateOrder{
		ProfileID:     body.ProfileID,
		BranchID:      body.BranchID,
		Date:          body.Date,
		CustomerPhone: body.CustomerPhone,
		Description:   body.Description,
		Status:        body.Status,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	httputils.WriteJSON(w, 201, body)
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
