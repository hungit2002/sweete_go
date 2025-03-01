package post

import (
	"context"
	log "github.com/sirupsen/logrus"
	"sweete/user_service/internal/domain/models"
	"sweete/user_service/internal/dto"
)

func (u *UseCase) GetPostsByParam(ctx context.Context, param dto.GetPostsByParamInputDTO) ([]*models.Post, error) {
	posts, err := u.postRepository.FindByParam(ctx, models.QueryPostParam{
		IDs:     param.IDs,
		Checkin: param.Checkin,
		Content: param.Content,
		UserID:  param.UserID,
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return posts, nil
}
