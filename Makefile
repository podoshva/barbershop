POSTGRES_DSN ?= postgres://admin:1234@localhost:5432/barbershop?sslmode=disable

migrate-up:
	goose -dir internal/adapter/postgres/migrations postgres "$(POSTGRES_DSN)" up


