package usecases

import (
	"context"
	"sweete/user_service/internal/dto"
)

type IAuthUseCase interface {
	Login(ctx context.Context, email string, phone string, password string) (*dto.UserResponse, string, error)
	Register(ctx context.Context, firstName string, surname string, dateOfBirth string, gender int, phone string, email string, password string) (*dto.UserResponse, string, error)
}
