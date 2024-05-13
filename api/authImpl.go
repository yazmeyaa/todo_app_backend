package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/yazmeyaa/todo_app_backend/data/request"
	"github.com/yazmeyaa/todo_app_backend/data/response"
	"github.com/yazmeyaa/todo_app_backend/models"
	"github.com/yazmeyaa/todo_app_backend/services"
)

type AuthControllerImpl struct {
	authService services.AuthService
	validate    *validator.Validate
}

// Login implements AuthController.
func (a *AuthControllerImpl) Login(ctx *gin.Context) {
	reqBody := request.LoginRequest{}

	jsonErr := ctx.ShouldBindJSON(&reqBody)
	if jsonErr != nil {
		ctx.JSON(400, response.NewApiErrorResponse("Bad JSON"))
		return
	}

	validationErr := a.validate.Struct(reqBody)
	if validationErr != nil {
		ctx.JSON(400, response.NewApiErrorResponse("Not valid JSON"))
		return
	}

	creds := services.Credentails{
		Username: &reqBody.Username,
		Email:    &reqBody.Email,
		Password: reqBody.Password,
	}

	token, loginErr := a.authService.Login(creds)
	if loginErr != nil {
		ctx.JSON(400, response.NewApiErrorResponse("Wrong username or password"))
		return
	}

	ctx.Header("auth", fmt.Sprintf("Bearer %s", token))
	ctx.Status(204)
}

// Register implements AuthController.
func (a *AuthControllerImpl) Register(ctx *gin.Context) {
	reqBody := request.RegisterRequest{}
	jsonError := ctx.ShouldBindJSON(&reqBody)
	if jsonError != nil {
		ctx.JSON(400, response.NewApiErrorResponse("Bad JSON"))
		return
	}

	validationErr := a.validate.Struct(reqBody)
	if validationErr != nil {
		ctx.JSON(400, response.NewApiErrorResponse("Not valid JSON"))
		return
	}

	newUser := models.User{
		Name:     reqBody.Name,
		Username: reqBody.Username,
		Email:    reqBody.Email,
		Password: reqBody.Password,
	}

	registerErr := a.authService.Register(&newUser)
	if registerErr != nil {
		ctx.JSON(400, response.NewApiErrorResponse(registerErr.Error()))
		return
	}

	ctx.Status(204)
}

func NewAuthController(authService services.AuthService, validator *validator.Validate) AuthController {
	return &AuthControllerImpl{
		authService: authService,
		validate:    validator,
	}
}
