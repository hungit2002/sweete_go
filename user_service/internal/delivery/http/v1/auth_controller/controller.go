package auth_controller

import (
	"sweete/user_service/internal/domain/usecases"
)

type Controller struct {
	authUseCase usecases.IAuthUseCase
}

func New(
	authUseCase usecases.IAuthUseCase,
) *Controller {
	return &Controller{
		authUseCase: authUseCase,
	}
}
