package friend

import "sweete/user_service/internal/domain/repositories"

type UseCase struct {
	friendRepository repositories.IFriendRepository
}

func New(
	friendRepository repositories.IFriendRepository,
) *UseCase {
	return &UseCase{
		friendRepository: friendRepository,
	}
}
