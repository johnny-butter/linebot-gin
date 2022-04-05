package cache

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	ctx context.Context
	rdb *redis.Client
}

func NewRedisClient() *RedisClient {
	rdbOpt, _ := redis.ParseURL(os.Getenv("CACHE_LOCATION"))

	return &RedisClient{
		ctx: context.Background(),
		rdb: redis.NewClient(rdbOpt),
	}
}

func (r *RedisClient) Get(key string) string {
	val, err := r.rdb.Get(r.ctx, key).Result()

	if err == redis.Nil {
		return ""
	}
	if err != nil {
		log.Println(err)
	}

	return val
}

func (r *RedisClient) Set(key string, value string, settings Settings) error {
	err := r.rdb.Set(r.ctx, key, value, settings.Ttl).Err()

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
