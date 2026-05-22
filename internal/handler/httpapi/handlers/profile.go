package handlers

import (
	"encoding/json"
	"main/internal/dto"
	"main/internal/handler/httputils"
	"net/http"
)

func (h *Handlers) CreateProfile(w http.ResponseWriter, r *http.Request) {
	var body struct {
		BranchID int64  `json:"branch_id,omitempty"`
		FullName string `json:"full_name,omitempty"`
		Login    string `json:"login,omitempty"`
		Password string `json:"password,omitempty"`
		Role     string `json:"role,omitempty"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := h.ProfileService.Create(r.Context(), dto.CreateProfile{
		BranchID: body.BranchID,
		FullName: body.FullName,
		Login:    body.Login,
		Password: body.Password,
		Role:     body.Role,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	httputils.WriteJSON(w, 201, body)
}

func (h *Handlers) DeleteProfile(w http.ResponseWriter, r *http.Request) {
	id, err := httputils.GetID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := h.ProfileService.Delete(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	httputils.WriteJSON(w, 204, nil)
}

func (h *Handlers) GetProfile(w http.ResponseWriter, r *http.Request) {
	id, err := httputils.GetID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	profile, err := h.ProfileService.Get(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	httputils.WriteJSON(w, 200, profile)
}
