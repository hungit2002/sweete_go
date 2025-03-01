package usecases

import (
	"context"
	"sweete/user_service/internal/domain/models"
	"sweete/user_service/internal/dto"
)

type IPostUseCase interface {
	Create(ctx context.Context, post *models.Post) error
	FindByParam(ctx context.Context, param models.QueryPostParam) ([]*models.Post, error)
	FirstByParam(ctx context.Context, param models.QueryPostParam) (*models.Post, error)
	GetPostsByParam(ctx context.Context, param dto.GetPostsByParamInputDTO) ([]*models.Post, error)
}
