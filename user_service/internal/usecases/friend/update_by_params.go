package friend

import (
	"context"
	log "github.com/sirupsen/logrus"
	"sweete/user_service/internal/domain/models"
)

func (u *UseCase) UpdateByParams(ctx context.Context, param models.QueryFriendParam, dataUpdate interface{}) error {
	err := u.friendRepository.UpdateByParams(ctx, param, dataUpdate)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
