package post

import (
	"context"
	log "github.com/sirupsen/logrus"
	"sweete/user_service/internal/domain/models"
)

func (u *UseCase) FirstByParam(ctx context.Context, param models.QueryPostParam) (*models.Post, error) {
	post, err := u.postRepository.FirstByParam(ctx, param)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return post, nil
}
