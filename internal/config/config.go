package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	HTTPPort             string
	DatabaseURL          string
	JWTSecret            string
	FilePathForSaveTasks string

	CookiesMaxAge   int
	CookiesPath     string
	CookiesDomain   string
	CookiesSecure   bool
	CookiesHttpOnly bool
}

func MustLoad() *Config {
	cfg := &Config{
		HTTPPort:             getEnv("HTTP_PORT", "8080"),
		DatabaseURL:          getEnv("DATABASE_URL", ""),
		JWTSecret:            getEnv("JWT_SECRET", ""),
		FilePathForSaveTasks: getEnv("FILE_PATH_FOR_SAVE_TASKS", "storage/task.json"),

		// HTTP-Only Cookies settings
		CookiesMaxAge:   getEnvInt("COOKIES_MAX_AGE", 604800),
		CookiesPath:     getEnv("COOKIES_PATH", "/"),
		CookiesDomain:   getEnv("COOKIES_DOMAIN", "localhost"),
		CookiesSecure:   getEnvBool("COOKIES_SECURE", false),
		CookiesHttpOnly: getEnvBool("COOKIES_HTTP_ONLY", true),
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

func getEnvBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	v, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}

	return v
}

func getEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	v, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return v
}
