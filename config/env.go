package config

import (
	"log"

	"github.com/caarlos0/env/v9"
)

type Environment struct {
	JWTSecret         string `env:"JWT_SECRET,required"`
	JWTExpirationDays int    `env:"JWT_EXPITATION_DAYS,required"`
	DatabaseName      string `env:"DATABASE_NAME,required"`
	DatabaseUser      string `env:"DATABASE_USER,required"`
	DatabasePassword  string `env:"DATABASE_PASSWORD,required"`
	DatabaseHost      string `env:"DATABASE_HOST,required"`
	DatabasePort      string `env:"DATABASE_PORT,required"`
	DatabaseSSLMode   string `env:"DATABASE_SSLMODE,required"`
	DatabaseLogMode   string `env:"DATABASE_LOGMODE,required"`
}

var environ Environment

func init() {
	if err := env.Parse(&environ); err != nil {
		log.Fatalf("Error al leer las variables de entorno: %v", err)
	}
}

func GetEnv() Environment {
	return environ
}
