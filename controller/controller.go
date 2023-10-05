package controller

import (
	"duck-cook-auth/usecase"
)

type Controller struct {
	authUseCase     usecase.AuthUseCase
	customerUseCase usecase.CustomerUseCase
}

func NewController(
	authUseCase usecase.AuthUseCase,
	customerUseCase usecase.CustomerUseCase,
) Controller {
	return Controller{
		authUseCase,
		customerUseCase,
	}
}
