package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	DbName string `env:"DB_NAME"`
	DbHost string `env:"DB_HOST"`
	DbUser string `env:"DB_USER"`
	DbPass string `env:"DB_PASS"`
	DbPort string `env:"DB_PORT"`
}

func (c *Config) DatabaseUrl() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		c.DbUser,
		c.DbPass,
		c.DbHost,
		c.DbPort,
		c.DbName,
	)
}

func New() (*Config, error) {
	conf, err := env.ParseAs[Config]()
	if err != nil {
		return nil, fmt.Errorf("failed to load the config: %w", err)
	}
	return &conf, nil
}
