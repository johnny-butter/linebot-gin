package cache

import (
	"time"
)

type Settings struct {
	Ttl time.Duration
}

type CacheClient interface {
	Get(key string) string
	Set(key string, value string, settings Settings) error
}

type CacheType int

const (
	Redis CacheType = iota
)

func NewCache(cacheType CacheType) CacheClient {
	switch cacheType {
	case Redis:
		return NewRedisClient()
	default:
		return NewRedisClient()
	}
}
