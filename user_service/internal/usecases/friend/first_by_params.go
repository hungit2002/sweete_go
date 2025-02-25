package friend

import (
	"context"
	"sweete/user_service/internal/domain/models"
)

func (u *UseCase) FirstByParams(ctx context.Context, param models.QueryFriendParam) (*models.Friend, error) {
	friend, err := u.friendRepository
}
