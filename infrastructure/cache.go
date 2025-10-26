package infrastructure

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	redisClient *redis.Client
}

func NewCache() *Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return &Cache{
		redisClient: rdb,
	}
}
func (rdb *Cache) Close() error {
	return rdb.redisClient.Close()
}

func (rdb *Cache) Set(ctx context.Context, key string, value string) error {
	expiration := 10 * time.Minute
	return rdb.redisClient.Set(ctx, key, value, expiration).Err()
}

func (rdb *Cache) Get(ctx context.Context, key string) (string, error) {
	return rdb.redisClient.Get(ctx, key).Result()
}
