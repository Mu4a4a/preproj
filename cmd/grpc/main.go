package main

import (
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log/slog"
	"net"
	"os"
	"preproj/cmd/di"
	"preproj/internal/config"
	"preproj/internal/handler/grpcapi"
	pbProduct "preproj/internal/handler/grpcapi/gen/product"
	pbUser "preproj/internal/handler/grpcapi/gen/user"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	if err := config.Init(); err != nil {
		slog.Error("failed to connect db", slog.Any("error", err))
		os.Exit(1)
	}

	services, cache, err := di.InitDependencies()
	if err != nil {
		slog.Error("failed to init dependencies", slog.Any("error", err))
		os.Exit(1)
	}
	ttl := viper.GetDuration("ttl.GRPC")
	srv := grpc.NewServer(grpc.UnaryInterceptor(grpcapi.CacheMiddleware(cache, ttl)))

	pbUser.RegisterUserServiceServer(srv, &grpcapi.UserService{Services: services})
	pbProduct.RegisterProductServiceServer(srv, &grpcapi.ProductService{Services: services})
	reflection.Register(srv)

	listener, err := net.Listen("tcp", viper.GetString("portGRPC"))
	if err != nil {
		slog.Error("failed to listen", slog.Any("error", err))
		os.Exit(1)
	}
	defer listener.Close()
	logger.Info("starting server", slog.Any("port", viper.GetString("portGRPC")))

	if err := srv.Serve(listener); err != nil {
		slog.Error("failed to start server", slog.Any("error", err))
		os.Exit(1)
	}
}
