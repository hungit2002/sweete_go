package repositories

import (
	"context"
	"sweete/user_service/internal/domain/models"
)

type IPostRepository interface {
	FindByParam(ctx context.Context, param models.QueryPostParam) ([]*models.Post, error)
	FirstByParam(ctx context.Context, param models.QueryPostParam) (*models.Post, error)
	Create(ctx context.Context, post *models.Post) error
}
