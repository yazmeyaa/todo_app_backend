package repository

import (
	"errors"
	"fmt"

	"github.com/yazmeyaa/todo_app_backend/models"
	"gorm.io/gorm"
)

type TaskRepositoryImpl struct {
	db *gorm.DB
}

func (r TaskRepositoryImpl) Create(task *models.Task) error {
	task.Status = models.STATUS_CREATED
	if err := r.db.Model(&models.Task{}).Create(&task).Error; err != nil {
		errMessage := fmt.Sprintf("Error while creating task: %s", err.Error())
		fmt.Print(errMessage)
		return errors.New(errMessage)
	} else {
		return nil
	}

}

func (r TaskRepositoryImpl) Update(task *models.Task) error {
	err := r.db.Model(&models.Task{}).Where("id = ?", task.ID).Updates(task).Error
	if err != nil {
		return errors.New("error while update task")
	} else {
		return nil
	}
}

func (r TaskRepositoryImpl) Delete(id uint) {
	r.db.Model(&models.Task{}).Where("id = ?", id).Delete(&models.Task{})
}

func (r TaskRepositoryImpl) FindById(id uint) (task models.Task, err error) {
	foundTask := models.Task{}

	result := r.db.Model(&models.Task{}).Where("id = ?", id).First(&foundTask)

	if result.Error != nil {
		return foundTask, errors.New(result.Error.Error())
	}

	return foundTask, nil
}

func (r TaskRepositoryImpl) GetList(options GetListOptions) (TaskList, error) {
	var results []models.Task
	var count int64

	query := r.db.Model(&models.Task{})
	query.Count(&count)
	query.Limit(int(options.Limit))
	query.Offset(int(options.Offset))
	query.Order("created_at desc")

	if options.status != nil {
		query.Where("status = ?", options.status)
	}

	if err := query.Find(&results).Error; err != nil {
		return TaskList{}, err
	}

	result := TaskList{
		Items: results,
		Count: uint(count),
	}

	return result, nil
}

func NewTaskRepository(database *gorm.DB) TaskRepository {
	return &TaskRepositoryImpl{database}
}
