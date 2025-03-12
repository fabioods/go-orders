package main

import (
	"log"

	"github.com/fabioods/go-orders/internal/handler"
	"github.com/fabioods/go-orders/internal/infra/rds"
	"github.com/fabioods/go-orders/internal/infra/s3"
	"github.com/fabioods/go-orders/internal/infra/webserver"
	"github.com/fabioods/go-orders/internal/usecase"
)

func main() {

	userRepositoryRDS := rds.NewUserRepositoryRDS()
	s3Repository := s3.NewUploadRepository()

	userUseCase := usecase.NewCreateUserUseCase(userRepositoryRDS)
	userAvartaUseCase := usecase.NewUserAvatarUseCase(userRepositoryRDS, s3Repository)

	userHandler := handler.NewUserHandler(userUseCase, userAvartaUseCase)

	webServer := webserver.NewWebServer(":8080")
	userHandler.AddUserHandler(webServer)

	log.Println("Server is running on port 8080")
	if err := webServer.Start(); err != nil {
		log.Fatal(err)
	}

}
