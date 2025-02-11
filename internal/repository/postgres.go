package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"preproj/internal/config"
	"time"
)

func NewPostgresDB() (*sql.DB, error) {
	cfg := config.LoadPostgresConfig()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode))
	if err != nil {
		slog.Error("failed to open database connection", slog.Any("error", err))
		return nil, errors.New("failed to open database connection")
	}
	err = db.PingContext(ctx)
	if err != nil {
		slog.Error("failed to ping database", slog.Any("error", err))
		_ = db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}
