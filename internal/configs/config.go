package config

import (
	"gem/pkg/env"
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	Addr      string
	JWTSecret string
	DB        struct {
		Addr         string
		MaxOpenConns int
		MaxIdleConns int
		MaxIdleTime  string
	}
}

func Load() (*Config, error) {

	err := godotenv.Load(".env.local")

	if err != nil {
		log.Fatal("❌ Failed to load env ", err)
	}

	cfg := &Config{}
	cfg.Addr = ":" + env.GetString("SERVER_ADDRESS")
	cfg.JWTSecret = env.GetString("JWT_SECRET")
	cfg.DB.Addr = env.GetString("DB_URL")

	cfg.DB.MaxOpenConns = 30
	cfg.DB.MaxIdleConns = 30
	cfg.DB.MaxIdleTime = "15m"

	return cfg, nil
}
