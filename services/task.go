package services

import (
	"github.com/yazmeyaa/todo_app_backend/data/response"
	"github.com/yazmeyaa/todo_app_backend/models"
	"github.com/yazmeyaa/todo_app_backend/repository"
)

type TaskService interface {
	Create(task *models.Task) error
	Update(task *models.Task) error
	Delete(id uint)
	FindById(id uint) (models.Task, error)
	GetList(options repository.GetListOptions) (repository.TaskList, error)
	PrepareTaskResponse(task *models.Task) response.TaskResponse
}
