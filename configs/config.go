package configs

import (
	"fmt"
	"os"
)

// DBConfig holds database connection parameters
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// Config holds all configuration for the application
type Config struct {
	Database DBConfig
	APIPort  string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	// Load database configuration
	dbConfig := DBConfig{
		Host:     getEnv("DATABASE_HOSTNAME", "localhost"),
		Port:     getEnv("DATABASE_PORT", "5432"),
		User:     getEnv("DATABASE_USERNAME", "postgres"),
		Password: getEnv("DATABASE_PASSWORD", ""),
		DBName:   getEnv("DATABASE_NAME", "go-db"),
	}

	// Validate required configurations
	if dbConfig.Password == "" {
		return nil, fmt.Errorf("database password is required")
	}

	// API configuration
	apiPort := getEnv("API_PORT", "8080")

	return &Config{
		Database: dbConfig,
		APIPort:  apiPort,
	}, nil
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}