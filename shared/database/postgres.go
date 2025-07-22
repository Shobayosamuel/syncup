package database

import (
    "database/sql"
    "fmt"
    "log"
    "time"

    _ "github.com/lib/pq"
)

type DB struct {
    *sql.DB
}

func NewPostgres(host string, port int, user, password, dbname, sslmode string) (*DB, error) {
    dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
        host, port, user, password, dbname, sslmode)

    db, err := sql.Open("postgres", dsn)
    if err != nil {
        return nil, fmt.Errorf("failed to open database: %w", err)
    }

    // Configure connection pool
    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(5)
    db.SetConnMaxLifetime(5 * time.Minute)

    // Test connection
    if err := db.Ping(); err != nil {
        return nil, fmt.Errorf("failed to ping database: %w", err)
    }

    log.Println("Successfully connected to PostgreSQL database")
    return &DB{db}, nil
}

func (db *DB) Close() error {
    return db.DB.Close()
}

// Health check method
func (db *DB) HealthCheck() error {
    return db.Ping()
}