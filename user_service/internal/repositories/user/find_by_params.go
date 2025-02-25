package user

import (
	"context"
	log "github.com/sirupsen/logrus"
	"sweete/user_service/internal/domain/models"
)

func (r *Repository) FindByParams(ctx context.Context, params models.QueryUserParam) ([]*models.User, error) {
	var results []*models.User
	db := r.GetDB(ctx).WithContext(ctx)
	query := PreloadUser(db, params.Preload)
	query = r.QueryUserByParam(query, params)

	if len(params.Select) > 0 {
		query = query.Select(query, params.Select)
	}

	if err := query.Find(&results).Error; err != nil {
		log.Error(err)
		return nil, err
	}
	return results, nil
}
