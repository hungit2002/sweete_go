package user

import (
	"context"
	log "github.com/sirupsen/logrus"
	"sweete/user_service/internal/domain/models"
	"sweete/user_service/internal/dto"
)

func (u *UseCase) GetUsersByParams(ctx context.Context, params dto.UserParamsDTO) ([]*models.User, error) {
	users, err := u.userRepository.FindByParams(ctx, models.QueryUserParam{Phone: params.Phone, Email: params.Email, FullName: params.FullName, IDs: params.IDs})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return users, nil
}
