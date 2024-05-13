package services

import (
	"errors"
	"log"

	"github.com/yazmeyaa/todo_app_backend/models"
	"golang.org/x/crypto/bcrypt"
)

var BCRYPT_COST = 6

type AuthServiceImpl struct {
	userSerivce UserService
	jwtService  JWTService
}

func NewAuthService(userService UserService, jwtService JWTService) AuthService {
	return &AuthServiceImpl{
		userSerivce: userService,
		jwtService:  jwtService,
	}
}

func (service AuthServiceImpl) Register(user *models.User) error {
	return service.userSerivce.Create(user)
}
func (service AuthServiceImpl) Login(creds Credentails) (token string, err error) {
	found, findErr := service.userSerivce.FindByUsername(*creds.Username)
	if findErr != nil {
		return "", findErr
	}

	compareError := bcrypt.CompareHashAndPassword([]byte(found.Password), []byte(creds.Password))

	if compareError != nil {
		log.Default().Println(compareError.Error())
		return "", errors.New("wrong password")
	}

	return service.jwtService.Sign(UserClaims{
		UserId: int(found.ID),
	})
}
