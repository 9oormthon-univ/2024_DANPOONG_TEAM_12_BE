package config

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	Env        string `env:"COURSEPICK_ENV" envDefault:"dev"`
	AppPort    string `env:"APP_PORT" envDefault:"8080"`
	DBHost     string `env:"DB_HOST" envDefault:"db"`
	DBPort     int    `env:"DB_PORT" envDefault:"3306"`
	DBUser     string `env:"DB_USER" envDefault:"coursepick"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"coursepick_pw"`
	DBName     string `env:"DB_NAME" envDefault:"coursepick_db"`
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
