package repositories

import (
	"context"
	"sweete/user_service/internal/domain/models"
)

type IFriendRepository interface {
	FindByParams(ctx context.Context, param models.QueryFriendParam) ([]*models.Friend, error)
	FirstByParams(ctx context.Context, param models.QueryFriendParam) (*models.Friend, error)
}
