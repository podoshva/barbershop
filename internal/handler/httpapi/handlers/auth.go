package handlers

import (
	"encoding/json"
	"main/internal/dto"
	"net/http"
)

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	var dto dto.LoginIn
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	token, err := h.AuthService.Login(r.Context(), dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Set-Cookie", *token)
}
