package services

import (
	"github.com/yazmeyaa/todo_app_backend/models"
	"github.com/yazmeyaa/todo_app_backend/repository"
)

type UserServiceImpl struct {
	Repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &UserServiceImpl{
		Repo: repo,
	}
}

func (service UserServiceImpl) Create(user *models.User) error {
	return service.Repo.Create(user)
}
func (service UserServiceImpl) Delete(userId int) {
	service.Repo.Delete(userId)
}
func (service UserServiceImpl) FindById(userId int) (*models.User, error) {
	return service.Repo.FindById(userId)
}
func (service UserServiceImpl) FindByUsername(username string) (*models.User, error) {
	return service.Repo.FindByUsername(username)
}
