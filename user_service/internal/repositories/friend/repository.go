package friend

import (
	"gorm.io/gorm"
	"sweete/user_service/internal/domain/models"
	"sweete/user_service/internal/repositories/base"
)

type Repository struct {
	base.BaseRepository
}

func New() *Repository {
	return &Repository{}
}

func PreloadFriend(db *gorm.DB, preload *models.PreloadFriend) *gorm.DB {
	return db
}

func QueryFriendByParams(query *gorm.DB, param models.QueryFriendParam) *gorm.DB {
	if param.FriendID != 0 {
		query = query.Where("friend_id = ?", param.FriendID)
	}
	if param.UserID != 0 {
		query = query.Where("user_id = ?", param.UserID)
	}
	if param.Status != 0 {
		query = query.Where("status = ?", param.Status)
	}
	if param.DeletedAt {
		query = query.Where("deleted_at is null")
	}
	return query
}
