// Package handlers
package handlers

import (
	"encoding/json"
	"main/internal/dto"
	"main/internal/handler/httputils"
	"net/http"
)

func (h *Handlers) CreateBranch(w http.ResponseWriter, r *http.Request) {
	var dto dto.CreateBranch
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := h.BranchService.Create(r.Context(), dto); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	httputils.WriteJSON(w, 201, dto)
}

func (h *Handlers) DeleteBranch(w http.ResponseWriter, r *http.Request) {
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

func (h *Handlers) GetBranch(w http.ResponseWriter, r *http.Request) {
	id, err := httputils.GetID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	branch, err := h.BranchService.Get(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	httputils.WriteJSON(w, 200, branch)
}
