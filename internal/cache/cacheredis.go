package cache

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"log/slog"
	"preproj/internal/config"
	"time"
)

type RedisClient struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisCache(ctx context.Context) RedisClient {
	cfg := config.LoadRedisConfig()

	client := redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
		Password:     cfg.Password,
		DB:           cfg.DB,
		Username:     cfg.User,
		MaxRetries:   cfg.MaxRetries,
		DialTimeout:  cfg.DialTimeout,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatal(err)
	}

	return RedisClient{
		client: client,
		ctx:    ctx,
	}
}

func (r RedisClient) Set(key string, value string, ttl time.Duration) error {
	slog.Debug("Setting key", slog.String("key", key), slog.String("value", value), slog.Duration("TTL", ttl))
	err := r.client.Set(r.ctx, key, value, ttl).Err()
	if err != nil {
		slog.Error("Error setting cache", slog.Any("error", err), slog.String("key", key))
		return fmt.Errorf("failed to set cache: %w", err)
	}
	slog.Debug("Successfully set cache", slog.String("key", key))
	return nil
}

func (r RedisClient) Get(key string) (string, error) {
	slog.Debug("Getting key", slog.String("key", key))
	value, err := r.client.Get(r.ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			slog.Debug("Key not found", slog.String("key", key))
			return "", ErrKeyNotFound
		}
		return "", fmt.Errorf("failed to get key from cache: %w", err)
	}
	slog.Debug("Successfully retrieved cache", slog.String(key, value))
	return value, nil
}
