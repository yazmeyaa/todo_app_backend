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
	cfg := config.NewAppConfig()
	jwtService := services.NewJwtService(cfg)
	validator := validator.New(validator.WithRequiredStructEnabled())

	db := config.GetDBConfig()
	db.Table("tasks").AutoMigrate(&models.Task{})
	db.Table("users").AutoMigrate(&models.User{})

	taskRepo := repository.NewTaskRepository(db)
	taskService := services.NewTaskService(taskRepo)

	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	authService := services.NewAuthService(userService, jwtService)

	gin := api.NewRouter(&api.RouterControllers{
		Tasks: api.NewTaskController(taskService, validator),
		Users: api.NewUserController(userService, validator, jwtService),
		Auth:  api.NewAuthController(authService, validator),
	}, validator, jwtService)

	gin.Run()
}
