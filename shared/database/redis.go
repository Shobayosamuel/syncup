package database

import (
    "context"
    "fmt"
    "time"

    "github.com/go-redis/redis/v8"
)

type Redis struct {
    *redis.Client
}

func NewRedis(host string, port int, password string, db int) (*Redis, error) {
    client := redis.NewClient(&redis.Options{
        Addr:         fmt.Sprintf("%s:%d", host, port),
        Password:     password,
        DB:           db,
        DialTimeout:  10 * time.Second,
        ReadTimeout:  30 * time.Second,
        WriteTimeout: 30 * time.Second,
        PoolSize:     10,
        PoolTimeout:  30 * time.Second,
    })

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := client.Ping(ctx).Err(); err != nil {
        return nil, fmt.Errorf("failed to connect to Redis: %w", err)
    }

    return &Redis{client}, nil
}

func (r *Redis) HealthCheck() error {
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()
    return r.Ping(ctx).Err()
}