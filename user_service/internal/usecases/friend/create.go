package friend

import (
	"context"
	log "github.com/sirupsen/logrus"
	"sweete/user_service/internal/domain/models"
)

func (u *UseCase) Create(ctx context.Context, friend *models.Friend) error {
	err := u.friendRepository.Create(ctx, friend)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
