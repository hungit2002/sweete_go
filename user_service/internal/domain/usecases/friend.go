package usecases

import (
	"context"
	"sweete/user_service/internal/domain/models"
	"sweete/user_service/internal/dto"
)

type IFriendUseCase interface {
	FindByParams(ctx context.Context, param models.QueryFriendParam) ([]*models.Friend, error)
	FirstByParams(ctx context.Context, param models.QueryFriendParam) (*models.Friend, error)
	Create(ctx context.Context, friend *models.Friend) error
	UpdateByParams(ctx context.Context, param models.QueryFriendParam, dataUpdate interface{}) error
	GetFriendByParam(ctx context.Context, param dto.GetFriendByParamInputDTO) ([]*models.User, error)
}
