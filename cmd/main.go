package main

import (
	"log"

	"github.com/fabioods/go-orders/internal/config"
	"github.com/fabioods/go-orders/internal/handler"
	"github.com/fabioods/go-orders/internal/infra/rds"
	"github.com/fabioods/go-orders/internal/infra/s3"
	"github.com/fabioods/go-orders/internal/infra/webserver"
	"github.com/fabioods/go-orders/internal/usecase"
)

func main() {

	cfg := config.LoadConfig()
	postgresConnection := rds.NewPostgresConnection(cfg)

	userRepositoryRDS := rds.NewUserRepositoryRDS(postgresConnection)
	s3Repository := s3.NewUploadRepository(cfg)
	orderRepositoryRDS := rds.NewOrderRepositoryRDS(postgresConnection)

	userUseCase := usecase.NewCreateUserUseCase(userRepositoryRDS)
	userAvartaUseCase := usecase.NewUserAvatarUseCase(userRepositoryRDS, s3Repository)

	orderUsecase := usecase.NewCreateOrderUseCase(userRepositoryRDS, orderRepositoryRDS)

	orderHandler := handler.NewOrderHandler(orderUsecase)
	userHandler := handler.NewUserHandler(userUseCase, userAvartaUseCase)

	webServer := webserver.NewWebServer(":8080")
	userHandler.AddUserHandler(webServer)
	orderHandler.AddOrderHandler(webServer)

	log.Println("Server is running on port 8080")
	if err := webServer.Start(); err != nil {
		log.Fatal(err)
	}

}
