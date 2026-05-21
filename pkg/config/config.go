package config

type Config struct {
	PostgresDsn string
}

func NewConfig() *Config {
	return &Config{
		"postgres://admin:1234@localhost:5432/barbershop?sslmode=disable",
	}
}
