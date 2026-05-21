package main

import (
	"context"
	"log"
	"main/internal/adapter/postgres"
	"main/internal/adapter/postgres/repos"
	"main/internal/app/branch"
	"main/internal/handler/httpapi"
	"main/internal/handler/httpapi/handlers"
	"main/pkg/config"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	cfg := config.NewConfig()
	pool, err := postgres.NewPool(ctx, cfg.PostgresDsn)
	if err != nil {
		panic(err)
	}
	branchRepo := repos.NewBranchRepo(pool)
	branchService := branch.NewBranchService(branchRepo)
	router := httpapi.NewRouter(httpapi.RouterDeps{
		Handlers: handlers.Handlers{
			BranchService: branchService,
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
