package user

import (
	"context"
	log "github.com/sirupsen/logrus"
	"sweete/user_service/internal/domain/models"
)

func (u *UseCase) GetUserDetail(ctx context.Context, id int) (*models.User, error) {
	user, err := u.userRepository.FirstByParams(ctx, models.QueryUserParam{ID: id, DeletedAt: true})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return user, nil
}
