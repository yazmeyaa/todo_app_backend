package services

import (
	"github.com/yazmeyaa/todo_app_backend/data/response"
	"github.com/yazmeyaa/todo_app_backend/models"
	"github.com/yazmeyaa/todo_app_backend/repository"
)

type TaskServiceImpl struct {
	repo repository.TaskRepository
}

func (service TaskServiceImpl) Create(task *models.Task) error {
	return service.repo.Create(task)
}
func (service TaskServiceImpl) Update(task *models.Task) error {
	return service.repo.Update(task)
}
func (service TaskServiceImpl) Delete(id uint) {
	service.repo.Delete(id)
}
func (service TaskServiceImpl) FindById(id uint) (models.Task, error) {
	return service.repo.FindById(id)

}
func (service TaskServiceImpl) GetList(options repository.GetListOptions) (repository.TaskList, error) {
	return service.repo.GetList(options)
}

func (service TaskServiceImpl) PrepareTaskResponse(task *models.Task) response.TaskResponse {
	response := response.TaskResponse{
		ID:        int(task.ID),
		DeletedAt: &task.DeletedAt.Time,
		CreatedAt: &task.CreatedAt,
		UpdatedAt: &task.UpdatedAt,
		Name:      task.Name,
		Status:    task.Status,
	}
	return response
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return TaskServiceImpl{
		repo: repo,
	}
}
