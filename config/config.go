package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	DATABASE_HOST     string
	DATABASE_PORT     int
	DATABASE_USER     string
	DATABASE_PASSWORD string
	DATABASE_NAME     string
}

func LoadConfig() *Config {
	port, err := strconv.Atoi(getEnv("DATABASE_PORT", "5432"))
	if err != nil {
		log.Fatalf("Invalid DATABASE_PORT value: %v", err)
	}

	return &Config{
		DATABASE_HOST:     getEnv("DATABASE_HOST", "localhost"),
		DATABASE_PORT:     port,
		DATABASE_USER:     getEnv("DATABASE_USER", "user"),
		DATABASE_PASSWORD: getEnv("DATABASE_PASSWORD", "password"),
		DATABASE_NAME:     getEnv("DATABASE_NAME", "dbname"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}