package config

import (
	"github.com/Netflix/go-env"
	_ "github.com/joho/godotenv/autoload"
)

var Parsed Config

func init() {
	if _, err := env.UnmarshalFromEnviron(&Parsed); err != nil {
		panic(err)
	}
}

type Config struct {
	Addr        string `env:"AUTH_ADDR,required=true"`
	Debug       bool   `env:"AUTH_DEBUG"`
	Origins     string `env:"AUTH_ORIGINS"`
	PostgresURL string `env:"AUTH_POSTGRES_URL,required=true"`
}
