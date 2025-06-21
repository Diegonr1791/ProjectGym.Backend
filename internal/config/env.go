package config

import (
	"os"
)

// Config maneja la configuración de la aplicación
type Config struct {
	DBHost               string
	DBPort               string
	DBUser               string
	DBPassword           string
	DBName               string
	ServerPort           string
	JWTSecret            string
	JWTExpirationMinutes string
}

// LoadConfig carga la configuración desde variables de entorno
func LoadConfig() *Config {
	return &Config{
		DBHost:               getEnv("DB_HOST", "localhost"),
		DBPort:               getEnv("DB_PORT", "5432"),
		DBUser:               getEnv("DB_USER", "postgres"),
		DBPassword:           getEnv("DB_PASSWORD", "admin"),
		DBName:               getEnv("DB_NAME", "gym-bro"),
		ServerPort:           getEnv("SERVER_PORT", "8080"),
		JWTSecret:            getEnv("JWT_SECRET", "supersecreto123"),
		JWTExpirationMinutes: getEnv("JWT_EXPIRATION_MINUTES", "60"),
	}
}

// getEnv obtiene una variable de entorno o devuelve un valor por defecto
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// GetJWTSecret implementa la interfaz JWTConfig
func (c *Config) GetJWTSecret() string {
	return c.JWTSecret
}

// GetJWTExpirationMinutes implementa la interfaz JWTConfig
func (c *Config) GetJWTExpirationMinutes() string {
	return c.JWTExpirationMinutes
}
