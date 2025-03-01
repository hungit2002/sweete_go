package post

import (
	"context"
	log "github.com/sirupsen/logrus"
	"sweete/user_service/internal/domain/models"
)

func (u *UseCase) FindByParam(ctx context.Context, param models.QueryPostParam) ([]*models.Post, error) {
	posts, err := u.postRepository.FindByParam(ctx, param)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return posts, nil
}
