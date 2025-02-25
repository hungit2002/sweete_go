package usecases

import (
	"context"
	"sweete/user_service/internal/domain/models"
	"sweete/user_service/internal/dto"
)

type IUserUseCaseInterface interface {
	GetUsersByParams(ctx context.Context, param dto.UserParamsDTO) ([]*models.User, error)
	GetUserDetail(ctx context.Context, id int) (*models.User, error)
	InviteFriend(ctx context.Context, userID int, friendID int, status int) error
}
