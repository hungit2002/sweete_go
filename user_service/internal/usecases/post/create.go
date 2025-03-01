package post

import (
	"context"
	log "github.com/sirupsen/logrus"
	"sweete/user_service/internal/domain/models"
)

func (u *UseCase) Create(ctx context.Context, post *models.Post) error {
	err := u.postRepository.Create(ctx, post)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
