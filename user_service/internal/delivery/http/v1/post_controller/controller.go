package post_controller

import "sweete/user_service/internal/domain/usecases"

type Controller struct {
	postUseCase usecases.IPostUseCase
}

func New(
	postUseCase usecases.IPostUseCase,
) *Controller {
	return &Controller{
		postUseCase: postUseCase,
	}
}
