package dataful

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type Cache interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string, document interface{}) error
}

type Redis struct {
	c *redis.Client
}

func NewRedis(ctx context.Context, redisURI string, redisPassword string) (*Redis, error) {
	redis := redis.NewClient(&redis.Options{
		Addr:     redisURI,
		Password: redisPassword,
		DB:       0,
	})
	rx := redis.Ping(ctx)
	if rx.Err() != nil {
		return nil, rx.Err()
	}

	return &Redis{
		c: redis,
	}, nil
}

func (r *Redis) Get(ctx context.Context, key string, document interface{}) error {
	result := r.c.Get(ctx, key)
	if result.Err() != nil {
		return result.Err()
	}

	b, err := result.Bytes()
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, document)
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	status := r.c.Set(ctx, key, data, expiration)
	return status.Err()
}
