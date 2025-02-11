package rest

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"preproj/internal/cache"
	"preproj/internal/service"
	"time"
)

func cacheMiddleware(service service.Cache, ttl time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// generate hash-key
		cacheKey := c.Request.Method + ":" + c.Request.URL.String()

		// try to get data
		cachedData, err := service.Get(cacheKey)
		if errors.Is(err, cache.ErrKeyNotFound) {
			slog.Error("failed to get cache", slog.Any("error", err))
			return
		} else if cachedData != "" {
			// if try is successful return data
			slog.Debug("cache get successful")
			c.Data(http.StatusOK, "application/json", []byte(cachedData))
			c.Abort() // abort try
			return
		}

		// continue request
		c.Next()

		// save response if status 200 ok
		if c.Writer.Status() == http.StatusOK {
			body, exists := c.Get("responseBody")
			if exists {
				jsonData, err := json.Marshal(body)
				if err != nil {
					slog.Error("failed to serialize data", slog.Any("err", err))
					return
				}
				err = service.Set(cacheKey, string(jsonData), ttl)
				if err != nil {
					slog.Error("failed to set cache", slog.Any("err", err))
				}
			}
		}
	}
}
