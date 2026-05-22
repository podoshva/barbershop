// Package httpapi
package httpapi

import (
	"main/internal/handler/httpapi/handlers"
	"net/http"
)

type RouterDeps struct {
	Handlers handlers.Handlers
}

func NewRouter(deps RouterDeps) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /branches/", deps.Handlers.CreateBranch)
	mux.HandleFunc("DELETE /branches/{id}/", deps.Handlers.DeleteBranch)
	mux.HandleFunc("GET /branches/{id}/", deps.Handlers.GetBranch)
	mux.HandleFunc("GET /branches/", deps.Handlers.GetAllBranches)

	mux.HandleFunc("POST /profiles/", deps.Handlers.CreateProfile)
	mux.HandleFunc("DELETE /profiles/{id}/", deps.Handlers.CreateProfile)
	mux.HandleFunc("GET /profiles/{id}/", deps.Handlers.GetProfile)
	mux.HandleFunc("GET /profiles/", deps.Handlers.GetAllProfiles)

	mux.HandleFunc("POST /orders/", deps.Handlers.CreateOrder)
	mux.HandleFunc("DELETE /orders/{id}/", deps.Handlers.DeleteOrder)
	mux.HandleFunc("GET /orders/{id}/", deps.Handlers.GetOrder)
	mux.HandleFunc("GET /orders/", deps.Handlers.GetAllOrders)
	mux.HandleFunc("PATCH /orders/{id}", deps.Handlers.SetOrderStatus)
	return mux
}
