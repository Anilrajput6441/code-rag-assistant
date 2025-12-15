package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port   string
	ApiKey string
}

func Load() Config {
	_ = godotenv.Load()

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		panic("GEMINI_API_KEY missing")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	return Config{
		Port:   port,
		ApiKey: apiKey,
	}
}
