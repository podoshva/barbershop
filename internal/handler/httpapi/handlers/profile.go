package handlers

import (
	"encoding/json"
	"main/internal/adapter/postgres/repos"
	"main/internal/handler/httputils"
	"net/http"
	"strconv"
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
	if err := h.ProfileService.Create(r.Context(), repos.CreateProfile{
		BranchID: body.BranchID,
		FullName: body.FullName,
		Login:    body.Login,
		Password: body.Password,
		Role:     body.Role,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	httputils.WriteJSON(w, 200, body)
}

func (h *Handlers) DeleteProfile(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := h.ProfileService.Delete(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
