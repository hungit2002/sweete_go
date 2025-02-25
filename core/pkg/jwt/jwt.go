package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"sweete/user_service/config"
	"time"
)

type MyClaims struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	jwt.StandardClaims
}

type JWTService interface {
	GenerateToken(ID int, email string, phone string) (string, error)
	VerifyToken(tokenString string) (*MyClaims, error)
}
type jwtService struct {
	secretKey string
}

func (j *jwtService) VerifyToken(tokenString string) (*MyClaims, error) {
	res, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("not valid")
		}
		return []byte(j.secretKey), nil
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if !res.Valid {
		log.Error(errors.New("token is invalid"))
		return nil, errors.New("token is invalid")
	}
	claims, ok := res.Claims.(*MyClaims)
	if !ok {
		return nil, errors.New("payload not valid")
	}
	return claims, nil
}

func getSecretKey() string {
	secretKey := config.GetInstance().App.SecretKey
	if secretKey == "" {
		log.Fatal("JWT secret key can't be empty")
	}
	return secretKey
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
	}
}

func (j *jwtService) GenerateToken(ID int, email string, phone string) (string, error) {
	claims := MyClaims{
		ID:    ID,
		Email: email,
		Phone: phone,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 24 * 365).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		log.Error(err)
		return "", err
	}
	return tokenString, nil
}

func getToken(c *gin.Context) string {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		authHeader = c.Request.Header.Get("Token")
	}
	if authHeader == "" {
		if value, ok := c.GetQuery("Authorization"); ok {
			authHeader = value
		} else if value2, ok2 := c.GetQuery("Token"); ok2 {
			authHeader = value2
		} else if value3, ok3 := c.GetQuery("token"); ok3 {
			authHeader = value3
		}
	}
	return authHeader
}
