package password

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type PasswordService interface {
	HashPassword(password string) (string, error)
	CheckPassword(hashPasswords string, password string) bool
}

type passwordService struct {
}

func (p *passwordService) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error(err)
		return "", err
	}
	return string(hash), nil
}

func (p *passwordService) CheckPassword(hashPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}

func NewPasswordService() PasswordService {
	return &passwordService{}
}
