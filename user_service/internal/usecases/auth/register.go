package auth

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"sweete/user_service/internal/domain/models"
	"sweete/user_service/internal/dto"
	"time"
)

func (u *UseCase) Register(ctx context.Context, firstName string, surname string, dateOfBirth string, gender int, phone string, email string, password string) (*dto.UserResponse, string, error) {
	// check user
	_, err := u.userRepository.FirstByParams(ctx, models.QueryUserParam{
		EmailOrPhone: &models.EmailOrPhone{
			Email: email,
			Phone: phone,
		},
		DeletedAt: true,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			goto Next
		}
		log.Error(err)
		return nil, "", err
	}
Next:

	parsedDate, err := time.Parse("2006-01-02", dateOfBirth)
	if err != nil {
		log.Error(err)
		return nil, "", err
	}
	hashPassword, err := u.passwordService.HashPassword(password)
	if err != nil {
		log.Error(err)
		return nil, "", err
	}

	user := &models.User{
		Phone:     phone,
		Email:     email,
		FullName:  firstName + " " + surname,
		Gender:    gender,
		Dob:       parsedDate,
		Password:  hashPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = u.userRepository.Create(ctx, user)
	if err != nil {
		log.Error(err)
		return nil, "", err
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
