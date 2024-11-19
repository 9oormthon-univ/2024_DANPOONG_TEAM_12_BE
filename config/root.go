package config

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	Env        string `env:"COURSEPICK_ENV"`
	AppPort    string `env:"APP_PORT"`
	DBHost     string `env:"DB_HOST"`
	DBPort     int    `env:"DB_PORT"`
	DBUser     string `env:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD"`
	DBName     string `env:"DB_NAME"`
}

func SetConfig() (*Config, error) {

	cfg := &Config{}
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("env load failed")
	}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
