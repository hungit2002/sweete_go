package user

import (
	"context"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"sweete/user_service/internal/domain/models"
)

func (r *Repository) FirstByParams(ctx context.Context, params models.QueryUserParam) (*models.User, error) {
	var result models.User
	db := r.GetDB(ctx).WithContext(ctx)
	query := PreloadUser(db, params.Preload)
	query = r.QueryUserByParam(query, params)

	if len(params.Select) > 0 {
		query = query.Select(params.Select)
	}
	if err := query.First(&result).Error; err != nil {
		log.Error(err)
		return nil, err
	}
	return &result, nil
}

func PreloadUser(db *gorm.DB, preload *models.PreloadUser) *gorm.DB {
	return db
}
