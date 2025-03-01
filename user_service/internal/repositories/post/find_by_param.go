package post

import (
	"context"
	log "github.com/sirupsen/logrus"
	"sweete/user_service/internal/domain/models"
)

func (r *Repository) FindByParam(ctx context.Context, param models.QueryPostParam) ([]*models.Post, error) {
	var posts []*models.Post

	db := r.GetDB(ctx).WithContext(ctx)
	query := PreloadPost(db, param.Preload)
	query = QueryPostByParam(query, param)

	if len(param.Select) > 0 {
		query = query.Select(param.Select)
	}
	if err := query.Find(&posts).Error; err != nil {
		log.Error(err)
		return nil, err
	}
	return posts, nil
}
