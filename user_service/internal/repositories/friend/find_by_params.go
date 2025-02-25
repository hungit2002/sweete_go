package friend

import (
	"context"
	log "github.com/sirupsen/logrus"
	"sweete/user_service/internal/domain/models"
)

func (r *Repository) FindByParams(ctx context.Context, param models.QueryFriendParam) ([]*models.Friend, error) {
	results := make([]*models.Friend, 0)

	db := r.GetDB(ctx).WithContext(ctx)
	query := PreloadFriend(db, param.Preload)
	query = QueryFriendByParams(query, param)

	if len(param.Select) > 0 {
		query = query.Select(param.Select)
	}
	if err := query.Find(&results).Error; err != nil {
		log.Error(err)
		return nil, err
	}
	return results, nil
}
