package friend

import (
	"context"
	log "github.com/sirupsen/logrus"
	"sweete/user_service/internal/domain/models"
)

func (u *UseCase) FirstByParams(ctx context.Context, param models.QueryFriendParam) (*models.Friend, error) {
	friend, err := u.friendRepository.FirstByParams(ctx, param)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return friend, nil
}
