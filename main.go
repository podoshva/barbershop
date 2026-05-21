package main

import (
	"context"
	"main/internal/adapter/postgres"
	"main/pkg/config"
)

func main() {
	ctx := context.Background()
	cfg := config.NewConfig()
	pool, err := postgres.NewPool(ctx, cfg.PostgresDsn)
	if err != nil {
		panic(err)
	}
	
}
