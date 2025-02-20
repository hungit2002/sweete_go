package auth

import (
	"sweete/core/pkg/jwt"
	"sweete/core/pkg/password"
	"sweete/user_service/internal/domain/repositories"
)

type UseCase struct {
	userRepository  repositories.IUserRepository
	jwtService      jwt.JWTService
	passwordService password.PasswordService
}

func New(
	userRepository repositories.IUserRepository,
	jwtService jwt.JWTService,
	passwordService password.PasswordService,
) *UseCase {
	return &UseCase{
		userRepository:  userRepository,
		jwtService:      jwtService,
		passwordService: passwordService,
	}
}
