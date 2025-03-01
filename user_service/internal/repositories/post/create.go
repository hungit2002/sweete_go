package post

import (
	"context"
	log "github.com/sirupsen/logrus"
	"sweete/user_service/internal/domain/models"
)

func (r *Repository) Create(ctx context.Context, post *models.Post) error {
	db := r.GetDB(ctx).WithContext(ctx)
	result := db.Create(&post)
	if result.Error != nil {
		log.Error(result.Error)
		return result.Error
	}
	return nil
}
