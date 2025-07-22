package main

import (
    "log"
    "github.com/Shobayosamuel/syncup/services/auth-service/internal/config"
    "github.com/Shobayosamuel/syncup/shared/database"
)

func main() {
    cfg := config.Load()

    // Test PostgreSQL
    db, err := database.NewPostgres(
        cfg.Database.Host,
        cfg.Database.Port,
        cfg.Database.User,
        cfg.Database.Password,
        cfg.Database.DBName,
        cfg.Database.SSLMode,
    )
    if err != nil {
        log.Fatal("Failed to connect to PostgreSQL:", err)
    }
    defer db.Close()

    // Test Redis
    redis, err := database.NewRedis(
        cfg.Redis.Host,
        cfg.Redis.Port,
        cfg.Redis.Password,
        cfg.Redis.DB,
    )
    if err != nil {
        log.Fatal("Failed to connect to Redis:", err)
    }
    // Use the redis variable to avoid 'declared and not used' error
    _ = redis

    log.Println("All database connections successful!")
}