package config

import "os"

type Config struct {
	DBNAME string
}

// New returns a new Config struct
func New() *Config {
	return &Config{
		DBNAME: getEnv("DBNAME", ""),
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
