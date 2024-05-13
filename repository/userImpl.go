package repository

import (
	"errors"

	"github.com/yazmeyaa/todo_app_backend/models"
	"golang.org/x/crypto/bcrypt"
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
	if user.Username == nil && user.Email == nil {
		return errors.New("username or email required")
	}

	if user.Email != nil || user.Username != nil {
		if user.Email != nil {
			findQuery = findQuery.Where("email = ?", *user.Email)
			if user.Username != nil {
				findQuery = findQuery.Or("username = ?", *user.Username)
			}
		} else if user.Username != nil {
			findQuery = findQuery.Where("username = ?", *user.Username)
			if user.Email != nil {
				findQuery = findQuery.Or("email = ?", *user.Email)
			}
		}
	}

	findError := findQuery.First(&existUser).Error
	isExist := true

	// It means user found
	if errors.Is(findError, gorm.ErrRecordNotFound) {
		isExist = false
	}

	if isExist {
		return errors.New("user with given username or email already exist")
	}

	BCRYPT_COST := 6
	hash, hashErr := bcrypt.GenerateFromPassword([]byte(user.Password), BCRYPT_COST)
	if hashErr != nil {
		return errors.New("falied to hash password")
	}

	newUser := models.User{
		Name:     user.Name,
		Username: user.Username,
		Password: string(hash),
		Email:    user.Email,
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
