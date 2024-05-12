package services

import (
	"errors"

	"github.com/yazmeyaa/todo_app_backend/models"
)

type AuthServiceImpl struct {
	userSerivce UserService
}

func NewAuthService(userService UserService) AuthService {
	return &AuthServiceImpl{
		userSerivce: userService,
	}
}

func (service AuthServiceImpl) Register(user *models.User) error {
	return service.userSerivce.Create(user)
}
func (service AuthServiceImpl) Login(creds Credentails) (token string, err error) {
	return "", errors.New("not implemented")
}
