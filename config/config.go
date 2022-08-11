package config

import (
	"fmt"
	"os"
)

type Config struct {
	AppPort string `json:"appport"`
	DSN     string `json:"dsn"`
}

func (c *Config) LoadConfig() error {
	// Implement viper load config

	c.AppPort = fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	c.DSN = os.Getenv("DSN")
	return nil
}
