package user

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"sweete/user_service/internal/domain/models"
	"time"
)

func (u *UseCase) InviteFriend(ctx context.Context, userID int, friendID int, status int) error {
	friend, err := u.friendUseCase.FirstByParams(ctx, models.QueryFriendParam{UserID: userID, FriendID: friendID, DeletedAt: true})
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error(err)
			return err
		}
	}
	if friend != nil {
		if friend.DeletedAt != nil {
			dataUpdate := map[string]interface{}{
				"deleted_at": nil,
				"status":     status,
			}
			err := u.friendUseCase.UpdateByParams(ctx, models.QueryFriendParam{UserID: userID, FriendID: friendID}, dataUpdate)
			if err != nil {
				log.Error(err)
				return err
			}
			return nil
		} else {
			return errors.New("friend exists")
		}
	}

	friend = &models.Friend{
		UserID:    userID,
		FriendID:  friendID,
		Status:    status,
		DeletedAt: nil,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = u.friendUseCase.Create(ctx, friend)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
