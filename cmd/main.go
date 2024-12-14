package main

import (
	"log"
	PreProject0 "preproj"
	"preproj/internal/handler"
	"preproj/internal/repository"
	"preproj/internal/service"
)

func main() {

	// TODO: db
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(PreProject0.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error staring servr: %s", err.Error())
	}
}
