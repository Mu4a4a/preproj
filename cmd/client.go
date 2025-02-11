package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"log/slog"
	"os"
	pb "preproj/internal/handler/grpcapi/gen/product"
	"time"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	// Устанавливаем соединение с сервером
	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Не удалось подключиться к серверу: %v", err)
	}
	defer conn.Close()
	logger.Info("Запуск сервера на %s")

	userClient := pb.NewProductServiceClient(conn)
	logger.Debug("Дебаг")
	// Создаем контекст с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Запрашиваем всех пользователей
	res, err := userClient.GetAllProducts(ctx, &pb.GetProductsRequest{})
	if err != nil {
		log.Fatalf("Ошибка при запросе всех пользователей: %v", err)
	}
	log.Printf("Ответ от сервера (GetAllUsers): %v", res)
}
