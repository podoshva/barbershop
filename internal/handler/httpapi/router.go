// Package httpapi
package httpapi

import (
	"main/internal/handler/httpapi/handlers"
	"main/internal/handler/httpapi/middlewares"
	"net/http"
)

type RouterDeps struct {
	Handlers handlers.Handlers
}

func NewRouter(deps RouterDeps) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /branches/", middlewares.Auth(middlewares.Access(middlewares.AccessRoleAdmin, deps.Handlers.CreateBranch)))
	mux.HandleFunc("DELETE /branches/{id}/", middlewares.Auth(middlewares.Access(middlewares.AccessRoleAdmin, deps.Handlers.DeleteBranch))) // admin
	mux.HandleFunc("GET /branches/{id}/", middlewares.Auth(deps.Handlers.GetBranch))                                                        // all
	mux.HandleFunc("GET /branches/", deps.Handlers.GetAllBranches)                                                                          // all

	mux.HandleFunc("POST /profiles/", middlewares.Auth(middlewares.Access(middlewares.AccessRoleAdmin, deps.Handlers.CreateProfile)))        // admin
	mux.HandleFunc("DELETE /profiles/{id}/", middlewares.Auth(middlewares.Access(middlewares.AccessRoleAdmin, deps.Handlers.CreateProfile))) // admin
	mux.HandleFunc("GET /profiles/{id}/", middlewares.Auth(deps.Handlers.GetProfile))
	mux.HandleFunc("GET /profiles/", middlewares.Auth(deps.Handlers.GetAllProfiles))

	mux.HandleFunc("POST /auth/login/", deps.Handlers.Login)

	mux.HandleFunc("POST /orders/", deps.Handlers.CreateOrder)                          // all
	mux.HandleFunc("DELETE /orders/{id}/", middlewares.Auth(deps.Handlers.DeleteOrder)) // nothing

	mux.HandleFunc("GET /orders/{id}/", middlewares.Auth(deps.Handlers.GetOrder)) // barber and admin

	mux.HandleFunc("GET /orders/", middlewares.Auth(middlewares.Access(middlewares.AccessRoleAdmin, deps.Handlers.GetAllOrders))) // barber and admin

	mux.HandleFunc("PATCH /orders/{id}", middlewares.Auth(deps.Handlers.SetOrderStatus)) // barber and admin
	return mux
}

/*
страница заказа
- создать заказ:
- список филиалов
- список барберов (profiles с фильтром по филиалу)

админ панель:
- создать барбера
- посмотреть все заказы (с фильтрами)
- список барберов (profiles с фильтрами)

страница барбера:
- посмотреть свои заказы (с фильтрами)
- изменить статус заказа
- выбрать рабочие дни и время
*/
