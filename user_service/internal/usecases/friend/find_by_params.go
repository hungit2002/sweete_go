package friend

import (
	"context"
	log "github.com/sirupsen/logrus"
	"sweete/user_service/internal/domain/models"
)

func (u *UseCase) FindByParams(ctx context.Context, param models.QueryFriendParam) ([]*models.Friend, error) {
	friends, err := u.friendRepository.FindByParams(ctx, param)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return friends, nil
}
