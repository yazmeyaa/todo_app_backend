package services

import "github.com/yazmeyaa/todo_app_backend/models"

type UserService interface {
	Create(*models.User) error
	Delete(int)
	FindById(int) (*models.User, error)
	FindByUsername(string) (*models.User, error)
}
