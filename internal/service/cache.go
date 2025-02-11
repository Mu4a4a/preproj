package service

import (
	"time"
)

type Cache interface {
	Set(key string, value string, ttl time.Duration) error
	Get(key string) (string, error)
}
type CacheService struct {
	cache Cache
}

func NewCache(cache Cache) *CacheService {
	return &CacheService{
		cache: cache,
	}
}

func (c *CacheService) Set(key string, value string, ttl time.Duration) error {
	return c.cache.Set(key, value, ttl)
}

func (c *CacheService) Get(key string) (string, error) {
	return c.cache.Get(key)
}
