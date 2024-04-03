package config

import (
	"os"

	"github.com/joho/godotenv"
)

type DB struct {
    DatabaseURL string
}

func New() (*DB, error) {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			return nil, err
		}
	}

  config := &DB{
    DatabaseURL: os.Getenv("LOCAL_DB_CONNECTION"),
  }

	return config, nil
}
