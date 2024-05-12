package repository

import "github.com/yazmeyaa/todo_app_backend/models"

type UserRepository interface {
	Create(*models.User) error
	Delete(int)
	FindById(int) (*models.User, error)
	FindByUsername(string) (*models.User, error)
}
