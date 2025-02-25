package usecases

import (
	"context"
	"sweete/user_service/internal/domain/models"
)

type IFriendUseCase interface {
	FindByParams(ctx context.Context, param models.QueryFriendParam) ([]*models.Friend, error)
	FirstByParams(ctx context.Context, param models.QueryFriendParam) (*models.Friend, error)
	Create(ctx context.Context, friend *models.Friend) error
}
