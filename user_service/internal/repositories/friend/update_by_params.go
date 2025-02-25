package friend

import (
	"context"
	log "github.com/sirupsen/logrus"
	"sweete/user_service/internal/domain/models"
)

func (r *Repository) UpdateByParams(ctx context.Context, param models.QueryFriendParam, dataUpdate interface{}) error {
	db := r.GetDB(ctx).WithContext(ctx)
	query := QueryFriendByParams(db, param)
	if err := query.Unscoped().Updates(dataUpdate).Error; err != nil {
		log.Error(err)
		return err
	}
	return nil
}
