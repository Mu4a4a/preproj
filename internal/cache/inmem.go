package cache

import (
	"errors"
	"log/slog"
	"sync"
	"time"
)

type CacheItem struct {
	Value      string
	Expiration int64
}
type Cache struct {
	mutex sync.RWMutex
	data  map[string]CacheItem
}

func NewInMemCache() *Cache {
	return &Cache{
		data: make(map[string]CacheItem),
	}
}

var (
	ErrKeyNotFound = errors.New("cache key not found")
	ErrKeyExpired  = errors.New("cache key has expired")
	ErrKeyEmpty    = errors.New("cache key is empty")
)

func (c *Cache) Set(key string, value string, ttl time.Duration) error {
	if key == "" {
		slog.Error("Invalid cache key", slog.String("cache key", key))
		return errors.New("empty cache key")
	}
	if ttl <= 0 {
		slog.Error("Invalid TTL", slog.Duration("TTL", ttl))
		return errors.New("TTL is equal to zero")
	}
	expiration := time.Now().Add(ttl).Unix()
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data[key] = CacheItem{
		Value:      value,
		Expiration: expiration,
	}
	return nil
}

func (c *Cache) Get(key string) (string, error) {
	if key == "" {
		slog.Error("Invalid cache key", slog.String("cache key", key))
		return "", ErrKeyEmpty
	}
	c.mutex.RLock()
	item, ok := c.data[key]
	if !ok {
		return "", ErrKeyNotFound
	}
	c.mutex.RUnlock()
	if time.Now().Unix() > item.Expiration {
		c.mutex.Lock()
		delete(c.data, key)
		c.mutex.Unlock()
		return "", ErrKeyExpired
	}
	return item.Value, nil
}
