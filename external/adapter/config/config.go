package config

import (
	"os"
)

type DB struct {
    DatabaseURL string
}

func New() (*DB, error) {
  config := &DB{
    DatabaseURL: os.Getenv("DATABASE_URL"),
  }

	return config, nil
}
