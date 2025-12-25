package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// Config holds runtime configuration values.
type Config struct {
	Addr   string
	DBHost string
	DBPort int
	DBUser string
	DBPass string
	DBName string
	DBSSL  string
}

// Load reads environment variables and builds a Config.
func Load() (Config, error) {
	_ = godotenv.Load()

	port, err := strconv.Atoi(getEnv("DB_PORT", "5432"))
	if err != nil {
		return Config{}, fmt.Errorf("invalid DB_PORT: %w", err)
	}

	cfg := Config{
		Addr:   getEnv("APP_ADDR", ":8080"),
		DBHost: getEnv("DB_HOST", "localhost"),
		DBPort: port,
		DBUser: getEnv("DB_USER", "postgres"),
		DBPass: getEnv("DB_PASSWORD", "postgres"),
		DBName: getEnv("DB_NAME", "postgres"),
		DBSSL:  strings.ToLower(getEnv("DB_SSLMODE", "disable")),
	}

	return cfg, nil
}

// DSN builds a PostgreSQL connection string for pgx.
func (c Config) DSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		urlEscape(c.DBUser),
		urlEscape(c.DBPass),
		c.DBHost,
		c.DBPort,
		c.DBName,
		c.DBSSL,
	)
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func urlEscape(value string) string {
	// Minimal escape for characters that break URLs.
	replacer := strings.NewReplacer("@", "%40", ":", "%3A", "/", "%2F")
	return replacer.Replace(value)
}
