package user_controller

import "sweete/user_service/internal/domain/usecases"

type Controller struct {
	userUseCase usecases.IUserUseCaseInterface
}

func New(
	userUseCase usecases.IUserUseCaseInterface,
) *Controller {
	return &Controller{
		userUseCase: userUseCase,
	}
}
