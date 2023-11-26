package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func New() (Config, error) {
	godotenv.Load(".env")

	config := Config{
		Port: os.Getenv("PORT"),
	}

	return config, nil
}
