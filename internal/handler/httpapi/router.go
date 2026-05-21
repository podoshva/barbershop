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
	mux.HandleFunc("POST /branch", deps.Handlers.CreateBranch)
	return mux
}
