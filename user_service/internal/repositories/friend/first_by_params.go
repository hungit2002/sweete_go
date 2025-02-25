package friend

import (
	"context"
	log "github.com/sirupsen/logrus"
	"sweete/user_service/internal/domain/models"
)

func (r *Repository) FirstByParams(ctx context.Context, param models.QueryFriendParam) (*models.Friend, error) {
	friend := &models.Friend{}
	db := r.GetDB(ctx).WithContext(ctx)
	query := PreloadFriend(db, param.Preload)
	query = QueryFriendByParams(query, param)

	if len(param.Select) > 0 {
		query = query.Select(param.Select)
	}
	if err := query.First(friend).Error; err != nil {
		log.Error(err)
		return nil, err
	}
	return friend, nil
}
