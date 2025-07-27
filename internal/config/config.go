package config

import (
	"os"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBName string
	DBHost string
	DBPort string
	JWTSecret string
}

func checkEnv(value string) (string) {
	s := os.Getenv(value)
	if s == "" {
		panic("Нету переменной: " + value)
	}
	return s
}

func NewConfig() (*Config, error) {
		dbUser := checkEnv("DB_USER")
		dbPassword := checkEnv("DB_PASSWORD")
		dbName := checkEnv("DB_NAME")
		dbHost := checkEnv("DB_HOST")
		dbPort := checkEnv("DB_PORT")
	return &Config{
		DBUser: dbUser,
		DBPassword: dbPassword,
		DBName: dbName,
		DBHost: dbHost,
		DBPort: dbPort,
	}, nil
}