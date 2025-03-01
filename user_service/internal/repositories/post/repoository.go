package post

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

func PreloadPost(query *gorm.DB, preload *models.PreloadPost) *gorm.DB {
	return query
}

func QueryPostByParam(query *gorm.DB, param models.QueryPostParam) *gorm.DB {
	if len(param.IDs) > 0 {
		query = query.Where("id in ?", param.IDs)
	}
	if param.ID != 0 {
		query = query.Where("id = ?", param.ID)
	}
	if param.Status != 0 {
		query = query.Where("status = ?", param.Status)
	}
	if param.UserID != 0 {
		query = query.Where("user_id = ?", param.UserID)
	}
	if param.Feeling != 0 {
		query = query.Where("feeling = ?", param.Feeling)
	}
	if param.Content != "" {
		query = query.Where("content like ?", "%"+param.Content+"%")
	}
	if param.DeletedAt {
		query = query.Where("deleted_at is not NULL")
	}
	return query
}
