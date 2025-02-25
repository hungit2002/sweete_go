package friend

import (
	"context"
	log "github.com/sirupsen/logrus"
	"sweete/user_service/internal/domain/models"
)

func (r *Repository) Create(ctx context.Context, friend *models.Friend) error {
	db := r.GetDB(ctx).WithContext(ctx)
	if err := db.Create(friend).Error; err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}
