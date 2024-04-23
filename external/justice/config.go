package justice

import (
	"os"
)

type Config struct {
    HeavenUrl string
    JusticeLegalUrl string
}

func NewConfig() (*Config, error) {
  config := &Config{
    HeavenUrl: os.Getenv("HEAVEN_URL"),
    JusticeLegalUrl: os.Getenv("JUSTICE_LEGAL_URL"),
  }

  return config, nil
}
