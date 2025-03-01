package post

import (
	"context"
	log "github.com/sirupsen/logrus"
	"sweete/user_service/internal/domain/models"
)

func (r *Repository) FirstByParam(ctx context.Context, param models.QueryPostParam) (*models.Post, error) {
	var post *models.Post
	db := r.GetDB(ctx).WithContext(ctx)
	query := PreloadPost(db, param.Preload)
	query = QueryPostByParam(query, param)

	if len(param.Select) > 0 {
		query = query.Select(param.Select)
	}
	if err := query.First(&post).Error; err != nil {
		log.Error(err)
		return nil, err
	}

	return post, nil
}
