package auth

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"sweete/user_service/internal/domain/models"
	"sweete/user_service/internal/dto"
)

func (u *UseCase) Login(ctx context.Context, email string, phone string, password string) (*dto.UserResponse, string, error) {
	user, err := u.userRepository.FirstByParams(ctx, models.QueryUserParam{Email: email, Phone: phone, DeletedAt: true})
	if err != nil {
		log.Error(err)
		return nil, "", err
	}

	if checked := u.passwordService.CheckPassword(user.Password, password); !checked {
		return nil, "", errors.New("Password or username invalid")
	}

	accessToken, err := u.jwtService.GenerateToken(user.ID, user.Email, user.Phone)
	if err != nil {
		log.Error(err)
		return nil, "", err
	}

	userResponse := &dto.UserResponse{
		ID:       user.ID,
		FullName: user.FullName,
		Phone:    user.Phone,
		Email:    user.Email,
		Avatar:   user.Avatar,
	}
	return userResponse, accessToken, nil
}
