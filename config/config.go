package config

import (
	"github.com/caarlos0/env/v6"
	"fmt"
)

// Env - configure environment variables
type Env struct {
	DbName string `env:"DB_NAME"` 
	DbHost string `env:"DB_HOST"`
	DbUser string `env:"DB_USER"`
	DbPassword string `env:"DB_PASSWORD"`
}


// EnvSetup - configure environment variables
func EnvSetup() Env {
	cfg := Env{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}
	return cfg
}
