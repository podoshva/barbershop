package handlers

import (
	"encoding/json"
	"main/internal/dto"
	"main/internal/handler/httputils"
	"net/http"
)

func (h *Handlers) CreateProfile(w http.ResponseWriter, r *http.Request) {
	var dto dto.CreateProfile
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := h.ProfileService.Create(r.Context(), dto); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	httputils.WriteJSON(w, 201, dto)
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

func (h *Handlers) GetAllProfiles(w http.ResponseWriter, r *http.Request) {
	profiles, err := h.ProfileService.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	httputils.WriteJSON(w, 200, profiles)
}
