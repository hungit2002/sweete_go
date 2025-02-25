package repositories

import (
	"context"
	"sweete/user_service/internal/domain/models"
)

type IUserRepository interface {
	IBaseRepository
	FirstByParams(ctx context.Context, param models.QueryUserParam) (*models.User, error)
	FindByParams(ctx context.Context, param models.QueryUserParam) ([]*models.User, error)

	Create(ctx context.Context, user *models.User) error
}
