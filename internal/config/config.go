package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Port      string `env:"SERVER_PORT"`
	JWTSecret string `env:"JWT_SECRET"`
	DB        Database
}

type Database struct {
	Host     string `env:"DATABASE_HOST"`
	Port     string `env:"DATABASE_PORT"`
	User     string `env:"DATABASE_USER"`
	Password string `env:"DATABASE_PASSWORD"`
	Name     string `env:"DATABASE_NAME"`
}

func Load() (*Config, error) {
	var cfg Config

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
