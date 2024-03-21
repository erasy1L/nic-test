package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoUrl string
}

func LoadConfig() (Config, error) {
	err := godotenv.Load(".env")
	cfg := Config{
		MongoUrl: os.Getenv("MONGO_URL"),
	}

	if err != nil {
		return cfg, err
	}

	return cfg, err
}
