package main

import (
	"duck-cook-auth/api"
	"duck-cook-auth/controller"
	"duck-cook-auth/pkg/mongo"
	"duck-cook-auth/repository/customer_repository"
	"duck-cook-auth/usecase"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Server api.Server
}

func NewAppConfig() AppConfig {
	_ = godotenv.Load()

	mongoDb := mongo.Connect()

	repositoryCustomer := customer_repository.New(mongoDb)

	loginUseCase := usecase.NewAuthUserCase()
	customerUseCase := usecase.NewCustomerUseCase(repositoryCustomer)

	controller := controller.NewController(loginUseCase, customerUseCase)
	server := api.NewServer(controller)

	return AppConfig{
		Server: *server,
	}
}
