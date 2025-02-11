package di

import (
	"log/slog"
	"os"
	"preproj/internal/cache"
	"preproj/internal/repository"
	"preproj/internal/service"
)

func InitDependencies() (*service.Service, *service.CacheService, error) {
	db, err := repository.NewPostgresDB()
	if err != nil {
		slog.Error("failed to connect db", slog.Any("error", err))
		os.Exit(1)
	}

	//redis := cache.NewRedisCache(ctx)
	inmem := cache.NewInMemCache()

	cache := service.NewCache(inmem)
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	return services, cache, nil
}
