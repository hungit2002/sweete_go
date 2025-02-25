package user

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"sweete/user_service/internal/domain/models"
)

func (u *UseCase) InviteFriend(ctx context.Context, userID int, friendID int, status int) error {
	_, err := u.friendUseCase.FirstByParams(ctx, models.QueryFriendParam{UserID: userID, FriendID: friendID})
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error(err)
			return err
		}
	}
	
	return nil
}
