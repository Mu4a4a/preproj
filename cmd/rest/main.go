package main

import (
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log/slog"
	"os"
	"preproj"
	"preproj/cmd/di"
	_ "preproj/internal/cache"
	"preproj/internal/config"
	"preproj/internal/handler/rest"
)

func main() {
	if err := config.Init(); err != nil {
		slog.Error("failed to init config", slog.Any("error", err))
		os.Exit(1)
	}
	services, cache, err := di.InitDependencies()
	if err != nil {
		slog.Error("failed to init dependencies", slog.Any("error", err))
		os.Exit(1)
	}
	handlers := rest.NewHandler(services, cache)

	srv := new(preproj.Server)
	if err := srv.Run(viper.GetString("portHTTP"), handlers.InitRoutes()); err != nil {
		slog.Error("failed to start server", slog.Any("error", err))
		os.Exit(1)
	}
}
