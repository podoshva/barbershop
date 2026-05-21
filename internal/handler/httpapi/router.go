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

	mux.HandleFunc("POST /profiles/", deps.Handlers.CreateProfile)
	mux.HandleFunc("DELETE /profiles/{id}/", deps.Handlers.CreateProfile)

	mux.HandleFunc("POST /orders/", deps.Handlers.CreateOrder)
	mux.HandleFunc("DELETE /orders/{id}/", deps.Handlers.DeleteOrder)
	return mux
}
