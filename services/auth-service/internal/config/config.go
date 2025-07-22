// internal/config/config.go
package config

import (
    "os"
    "strconv"
)

type Config struct {
    Database DatabaseConfig
    Redis    RedisConfig
    JWT      JWTConfig
    Server   ServerConfig
}

type DatabaseConfig struct {
    Host     string
    Port     int
    User     string
    Password string
    DBName   string
    SSLMode  string
}

type RedisConfig struct {
    Host     string
    Port     int
    Password string
    DB       int
}

type JWTConfig struct {
    SecretKey      string
    ExpirationHour int
}

type ServerConfig struct {
    Port int
}

func Load() *Config {
    return &Config{
        Database: DatabaseConfig{
            Host:     getEnv("DB_HOST", "localhost"),
            Port:     getEnvInt("DB_PORT", 5433),
            User:     getEnv("DB_USER", "auth_user"),
            Password: getEnv("DB_PASSWORD", "auth_password"),
            DBName:   getEnv("DB_NAME", "dating_auth"),
            SSLMode:  getEnv("DB_SSL_MODE", "disable"),
        },
        Redis: RedisConfig{
            Host:     getEnv("REDIS_HOST", "localhost"),
            Port:     getEnvInt("REDIS_PORT", 6380),
            Password: getEnv("REDIS_PASSWORD", ""),
            DB:       getEnvInt("REDIS_DB", 0),
        },
        JWT: JWTConfig{
            SecretKey:      getEnv("JWT_SECRET", "your-super-secret-key"),
            ExpirationHour: getEnvInt("JWT_EXPIRATION_HOUR", 24),
        },
        Server: ServerConfig{
            Port: getEnvInt("SERVER_PORT", 8081),
        },
    }
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
    if value := os.Getenv(key); value != "" {
        if intValue, err := strconv.Atoi(value); err == nil {
            return intValue
        }
    }
    return defaultValue
}