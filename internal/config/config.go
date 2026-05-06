package config

import (
	"log"
	"os"
)

type Config struct {
	HTTPPort             string
	DatabaseURL          string
	JWTSecret            string
	FilePathForSaveTasks string
}

func MustLoad() *Config {
	cfg := &Config{
		HTTPPort:             getEnv("HTTP_PORT", "8080"),
		DatabaseURL:          getEnv("DATABASE_URL", ""),
		JWTSecret:            getEnv("JWT_SECRET", ""),
		FilePathForSaveTasks: getEnv("FILE_PATH_FOR_SAVE_TASKS", ""),
	}

	if cfg.DatabaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}
	if cfg.JWTSecret == "" {
		log.Fatal("JWT_SECRET is required")
	}

	return cfg
}

func getEnv(key, defaultValue string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultValue
}
