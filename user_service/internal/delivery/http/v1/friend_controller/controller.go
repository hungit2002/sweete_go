package friend_controller

import "sweete/user_service/internal/domain/usecases"

type Controller struct {
	friendUseCase usecases.IFriendUseCase
}

func New(friendUseCase usecases.IFriendUseCase) *Controller {
	return &Controller{friendUseCase: friendUseCase}
}
