package config

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
)

type Config struct {
	Port               int
	DatabaseURL        string
	CodAPIBaseURL      string
	CodSSOToken        string
	CORSAllowedOrigins string
	LogLevelStr        string
}

func Load() (*Config, error) {
	cfg := &Config{
		Port:               getEnvInt("PORT", 8080),
		DatabaseURL:        getEnv("DATABASE_URL", ""),
		CodAPIBaseURL:      getEnv("COD_API_BASE_URL", "https://my.callofduty.com/api/papi-client"),
		CodSSOToken:        getEnv("COD_SSO_TOKEN", ""),
		CORSAllowedOrigins: getEnv("CORS_ALLOWED_ORIGINS", "http://localhost:5173"),
		LogLevelStr:        getEnv("LOG_LEVEL", "info"),
	}

	if cfg.DatabaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is required")
	}

	return cfg, nil
}

func (c *Config) LogLevel() slog.Level {
	switch c.LogLevelStr {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return fallback
}
