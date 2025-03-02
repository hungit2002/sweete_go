package friend

import (
	"sweete/user_service/internal/domain/repositories"
)

type UseCase struct {
	friendRepository repositories.IFriendRepository
	userRepository   repositories.IUserRepository
}

func New(
	friendRepository repositories.IFriendRepository,
	userRepository repositories.IUserRepository,
) *UseCase {
	return &UseCase{
		friendRepository: friendRepository,
		userRepository:   userRepository,
	}
}
