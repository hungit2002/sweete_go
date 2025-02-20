package user

import (
	"context"
	log "github.com/sirupsen/logrus"
	"sweete/user_service/internal/domain/models"
)

func (r *Repository) Create(ctx context.Context, user *models.User) error {
	db := r.GetDB(ctx).WithContext(ctx)
	result := db.Create(&user)
	if result.Error != nil {
		log.Error(result.Error)
		return result.Error
	}
	return nil
}
