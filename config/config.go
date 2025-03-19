package config

import (
	"os"
)

type AIConfig struct {
	BaseURL string
	APIKey  string
	Model   string
}

func LoadAIConfig() *AIConfig {
	return &AIConfig{
		BaseURL: getEnv("OPENAI_BASE_URL", "http://localhost:11434/v1"),
		APIKey:  getEnv("OPENAI_API_KEY", "ollama"),
		Model:   getEnv("OPENAI_MODEL", "llama3.2"),
	}
}

type DBConfig struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
}

func LoadDBConfig() *DBConfig {
	return &DBConfig{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "1234"),
		DBName:     getEnv("DB_NAME", "mobilo_go_server"),
		DBPort:     getEnv("DB_PORT", "5432"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
