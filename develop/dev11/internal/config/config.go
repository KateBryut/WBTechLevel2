package config

import (
	"fmt"

	"github.com/caarlos0/env"
)

type Config struct {
	Port         int `env:"PORT" envDefault:"8080"`
	ReadTimeout  int `env:"R_TIMEOUT" envDefault:"10"`
	WriteTimeout int `env:"W_TIMEOUT" envDefault:"10"`
}

func NewConfig() (Config, error) {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		return cfg, err
	}
	fmt.Println("port", cfg.Port)

	return cfg, nil
}
