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

type UserControllerImpl struct {
	Service    services.UserService
	Validate   *validator.Validate
	JwtService services.JWTService
}

func NewUserController(service services.UserService, validate *validator.Validate, jwtService services.JWTService) UserController {
	return &UserControllerImpl{
		Service:    service,
		Validate:   validate,
		JwtService: jwtService,
	}
}

func (controller UserControllerImpl) Create(ctx *gin.Context) {
	reqBody := request.CreateUserRequest{}

	bindError := ctx.ShouldBindJSON(&reqBody)
	if bindError != nil {
		ctx.JSON(400, response.ApiErrorResponse{
			Error: "Failed to parse JSON body",
		})
	}

	validationError := controller.Validate.Struct(reqBody)
	if validationError != nil {
		ctx.JSON(400, response.ApiErrorResponse{
			Error: "Not valid JSON",
		})
		return
	}

	if reqBody.Username == "" {
		ctx.JSON(400, response.ApiErrorResponse{
			Error: "Username is required",
		})
		return
	}

	newUser := models.User{
		Name:     reqBody.Name,
		Username: &reqBody.Username,
		Email:    &reqBody.Email,
		Password: reqBody.Password,
	}

	error := controller.Service.Create(&newUser)

	if error != nil {
		ctx.JSON(400, response.ApiErrorResponse{
			Error: fmt.Sprintf("Failed to create new user: %s", error.Error()),
		})
		return
	}

	ctx.Status(204)
}
func (controller UserControllerImpl) Delete(*gin.Context) {}
func (controller UserControllerImpl) Find(*gin.Context)   {}
