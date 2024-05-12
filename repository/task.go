package repository

import (
	"github.com/yazmeyaa/todo_app_backend/models"
	"github.com/yazmeyaa/todo_app_backend/utils"
)

type GetListOptions struct {
	utils.ListOptions
	status *string
}

type TaskList struct {
	Items []models.Task `json:"items,omitempty"`
	Count uint          `json:"total"`
}

type TaskRepository interface {
	Create(task *models.Task) error
	Update(task *models.Task) error
	Delete(id uint)
	FindById(id uint) (models.Task, error)
	GetList(options GetListOptions) (TaskList, error)
}
