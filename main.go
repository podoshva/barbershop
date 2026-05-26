package main

import (
	"context"
	"log"
	"main/internal/adapter/postgres"
	"main/internal/adapter/postgres/repos"
	"main/internal/app/auth"
	"main/internal/app/branch"
	"main/internal/app/order"
	"main/internal/app/profile"
	"main/internal/handler/httpapi"
	"main/internal/handler/httpapi/handlers"
	"main/pkg"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	cfg := pkg.NewConfig()
	pool, err := postgres.NewPool(ctx, cfg.PostgresDsn)
	if err != nil {
		panic(err)
	}

	branchRepo := repos.NewBranchRepository(pool)
	branchService := branch.NewBranchService(branchRepo)

	profileRepo := repos.NewProfileRepository(pool)
	profileService := profile.NewProfileService(profileRepo)
	authService := auth.NewAuthService(profileRepo)

	orderRepo := repos.NewOrderRepository(pool)
	orderService := order.NewOrderService(orderRepo)

	router := httpapi.NewRouter(httpapi.RouterDeps{
		Handlers: handlers.Handlers{
			BranchService:  branchService,
			ProfileService: profileService,
			OrderService:   orderService,
			AuthService:    authService,
		},
	})

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
