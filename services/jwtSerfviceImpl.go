package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yazmeyaa/todo_app_backend/config"
)

type JWTServiceImpl struct {
	appConfig *config.AppConfig
}

func NewJwtService(appConfig *config.AppConfig) JWTService {
	return &JWTServiceImpl{
		appConfig: appConfig,
	}
}

func (jwtService JWTServiceImpl) Sign(payload UserClaims) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(24)).Unix()

	claims := jwt.MapClaims{
		"userId": payload.UserId,
		"exp":    exp,
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	tokenString, err := token.SignedString([]byte(jwtService.appConfig.JWT.Secret))
	if err != nil {
		return "", fmt.Errorf("failed to create JWT: %s", err.Error())
	}

	return tokenString, nil

}
func (jwtService JWTServiceImpl) Verify(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtService.appConfig.JWT.Secret), nil
	})

	if err != nil {
		return false
	}
	return true
}

func (jwtService JWTServiceImpl) Decode(token string) (*UserClaims, error) {
	userClaims := UserClaims{}
	_, err := jwt.ParseWithClaims(token, &userClaims, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtService.appConfig.JWT.Secret), nil
	})

	if err != nil {
		return &userClaims, errors.New("failed to decode token")
	}

	return &userClaims, nil
}
