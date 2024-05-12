package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/yazmeyaa/todo_app_backend/api"
	"github.com/yazmeyaa/todo_app_backend/config"
	"github.com/yazmeyaa/todo_app_backend/models"
	"github.com/yazmeyaa/todo_app_backend/repository"
	"github.com/yazmeyaa/todo_app_backend/services"
)

func main() {
	validator := validator.New(validator.WithRequiredStructEnabled())

	db := config.GetDBConfig()
	db.Table("tasks").AutoMigrate(&models.Task{})
	db.Table("users").AutoMigrate(&models.User{})

	taskRepo := repository.NewTaskRepository(db)
	taskService := services.NewTaskService(taskRepo)

	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	gin := api.NewRouter(&api.RouterControllers{
		Tasks: api.NewTaskController(taskService, validator),
		Users: api.NewUserController(userService, validator),
	}, validator)

	gin.Run()
}
