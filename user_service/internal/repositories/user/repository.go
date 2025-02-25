package user

import (
	"gorm.io/gorm"
	"sweete/user_service/internal/domain/models"
	"sweete/user_service/internal/domain/repositories"
	"sweete/user_service/internal/repositories/base"
)

type Repository struct {
	base.BaseRepository
}

func New() repositories.IUserRepository {
	return &Repository{}
}

func PreloadUser(query *gorm.DB, preload *models.PreloadUser) *gorm.DB {
	if preload != nil {
		if preload.Friends != nil {
			query = query.Preload("Friends", func(db *gorm.DB) *gorm.DB {
				db = db.Where("deleted_at IS NULL")

				if len(preload.Friends.Query.Select) > 0 {
					db = db.Select(preload.Friends.Query.Select)
				}
				return db
			})
		}
		if preload.FriendUsers != nil {
			query = query.Preload("Friends", func(db *gorm.DB) *gorm.DB {
				db = db.Where("deleted_at IS NULL")
				return db
			})
			query = query.Preload("Friends.Users", func(db *gorm.DB) *gorm.DB {
				db = db.Where("deleted_at IS NULL")
				if len(preload.FriendUsers.Query.Select) > 0 {
					db = db.Select(preload.FriendUsers.Query.Select)
				}
				return db
			})
		}
	}
	return query
}

func (r *Repository) QueryUserByParam(query *gorm.DB, params models.QueryUserParam) *gorm.DB {
	if len(params.IDs) > 0 {
		query = query.Where("users.id in ?", params.IDs)
	}
	if params.ID != 0 {
		query = query.Where("users.id = ?", params.ID)
	}
	if params.Phone != "" {
		query = query.Where("users.phone = ?", params.Phone)
	}
	if params.Email != "" {
		query = query.Where("users.email = ?", params.Email)
	}
	if params.FullName != "" {
		query = query.Where("users.full_name like ?", "%"+params.FullName+"%")
	}
	if params.DeletedAt {
		query = query.Where("users.deleted_at is null")
	}

	if params.EmailOrPhone != nil {
		query = query.Where(
			"users.email = ? OR users.phone = ?",
			params.EmailOrPhone.Email,
			params.EmailOrPhone.Phone,
		)
	}
	return query
}
