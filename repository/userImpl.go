package repository

import (
	"errors"

	"github.com/yazmeyaa/todo_app_backend/models"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (repository UserRepositoryImpl) Create(user *models.User) error {
	var existUser models.User

	findQuery := repository.DB.Model(&models.User{})
	findError := findQuery.Where("username = ? OR email = ?", user.Username, user.Email).Find(&existUser).Error

	// It means user found
	if findError == nil {
		if *existUser.Email != "" {
			return errors.New("user with given email already exist")
		}
		if *existUser.Username != "" {
			return errors.New("user with given username already exist")
		}
	}

	newUser := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
		Password: user.Password,
	}
	query := repository.DB.Model(&models.User{})
	query.Create(&newUser)

	return query.Error
}

func (repository UserRepositoryImpl) Delete(userId int) {
	query := repository.DB.Model(&models.User{})
	query.Delete("id = ?", userId)
}

func (repository UserRepositoryImpl) FindById(userId int) (*models.User, error) {
	query := repository.DB.Model(&models.User{})
	foundUser := models.User{}

	query.Where("id = ?", userId).First(&foundUser)

	return &foundUser, query.Error
}

func (repository UserRepositoryImpl) FindByUsername(username string) (*models.User, error) {
	query := repository.DB.Model(&models.User{})
	foundUser := models.User{}

	query.Where("username = ?", username).First(&foundUser)

	return &foundUser, query.Error
}
