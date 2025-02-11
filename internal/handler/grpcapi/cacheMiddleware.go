package grpcapi

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"google.golang.org/grpc"
	"log"
	"log/slog"
	"preproj/internal/cache"
	"preproj/internal/service"
	"time"
)

func CacheMiddleware(service service.Cache, ttl time.Duration) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// Generate hash-key
		cacheKey := generateCacheKey(info.FullMethod, req)
		// Try to get data from cache
		cachedData, err := service.Get(cacheKey)
		if errors.Is(err, cache.ErrKeyNotFound) {
			slog.Error("error get cache", slog.Any("error", err))
			return nil, err
		} else if len(cachedData) > 0 {
			// If successful, return cached data
			var response interface{}
			if err := json.Unmarshal([]byte(cachedData), &response); err != nil {
				slog.Error("error unmarshalling cached data", slog.Any("error", err))
				return nil, err
			}
			return response, nil
		}
		// Handle the request
		res, err := handler(ctx, req)
		if err != nil {
			return nil, err
		}
		// Save response to cache if no error
		if res != nil {
			jsonData, err := json.Marshal(res)
			if err != nil {
				slog.Error("error marshaling response", err)
			} else {
				if err := service.Set(cacheKey, string(jsonData), ttl); err != nil {
					slog.Error("error setting cache", slog.Any("error", err))
				}
			}
		}
		return res, nil
	}
}
func generateCacheKey(method string, req interface{}) string {
	jsonData, err := json.Marshal(req)
	if err != nil {
		log.Printf("Error marshaling request: %v", err)
		return method
	}
	hash := sha256.Sum256(jsonData)
	return method + ":" + hex.EncodeToString(hash[:])
}
