package user

import (
	"sweete/user_service/internal/domain/repositories"
	"sweete/user_service/internal/domain/usecases"
)

type UseCase struct {
	userRepository repositories.IUserRepository
	friendUseCase  usecases.IFriendUseCase
}

func New(
	userRepository repositories.IUserRepository,
	friendUseCase usecases.IFriendUseCase,
) *UseCase {
	return &UseCase{
		userRepository: userRepository,
		friendUseCase:  friendUseCase,
	}
}
