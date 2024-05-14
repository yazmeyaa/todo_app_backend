package services

import "github.com/yazmeyaa/todo_app_backend/models"

type Credentails struct {
	Username *string
	Email    *string
	Password string
}

type AuthService interface {
	Register(user *models.User) error
	Login(creds Credentails) (token string, user *models.User, err error)
}
