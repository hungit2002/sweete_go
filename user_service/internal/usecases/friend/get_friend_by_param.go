package friend

import (
	"context"
	log "github.com/sirupsen/logrus"
	"sweete/user_service/internal/domain/models"
	"sweete/user_service/internal/dto"
)

func (u *UseCase) GetFriendByParam(ctx context.Context, param dto.GetFriendByParamInputDTO) ([]*models.User, error) {
	friends, err := u.friendRepository.FindByParams(ctx, models.QueryFriendParam{
		UserID:    param.UserID,
		DeletedAt: true,
	})
	if err != nil {
		return nil, err
	}

	if len(friends) == 0 {
		return make([]*models.User, 0), nil
	}

	friendIDs := make([]int, 0, len(friends))
	for _, friend := range friends {
		friendIDs = append(friendIDs, friend.FriendID)
	}

	users, err := u.userRepository.FindByParams(ctx, models.QueryUserParam{IDs: friendIDs, FullName: param.FullName, DeletedAt: true})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return users, nil
}
