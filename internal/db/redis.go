package db

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"ming/pkg/config"
)

var Redis *redis.Client

func ConnectRedis(cfg *config.Config) error {
	Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password, // no password set
		DB:       cfg.Redis.DB,       // use default DB
	})

	_, err := Redis.Ping(context.Background()).Result()
	if err != nil {
		return fmt.Errorf("failed to connect to Redis: %w", err)
	}
	return nil
}
