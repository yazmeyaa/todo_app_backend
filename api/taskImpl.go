package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/yazmeyaa/todo_app_backend/data/request"
	"github.com/yazmeyaa/todo_app_backend/data/response"
	"github.com/yazmeyaa/todo_app_backend/models"
	"github.com/yazmeyaa/todo_app_backend/repository"
	"github.com/yazmeyaa/todo_app_backend/services"
	"github.com/yazmeyaa/todo_app_backend/utils"
)

type TaskControllerImpl struct {
	Service  services.TaskService
	Validate *validator.Validate
}

func (controller *TaskControllerImpl) GetById(ctx *gin.Context) {
	idParam, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(400, response.NewApiErrorResponse("Missed id param"))
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, response.NewApiErrorResponse("Unexprected value in param :id"))
		return
	}

	task, err := controller.Service.FindById(uint(id))

	if err != nil {
		ctx.JSON(400, response.NewApiErrorResponse(err.Error()))
		return
	}
	ctx.JSON(201, response.GetByIdResponse{
		Task: controller.Service.PrepareTaskResponse(&task),
	})
}

func (controller *TaskControllerImpl) GetList(ctx *gin.Context) {
	var defaultLimit int = 10
	var defaultOffset int = 0
	var queryError error
	limit, queryError := utils.GetIntQuery(ctx, "limit", &defaultLimit)
	if queryError != nil {
		return
	}
	offset, queryError := utils.GetIntQuery(ctx, "offset", &defaultOffset)
	if queryError != nil {
		return
	}

	list, listError := controller.Service.GetList(repository.GetListOptions{
		ListOptions: utils.ListOptions{
			Offset: offset,
			Limit:  limit,
		},
	})

	if listError != nil {
		ctx.JSON(400, response.ApiErrorResponse{
			Error: listError.Error(),
		})
		return
	}

	listResponse := make([]response.TaskResponse, len(list.Items))

	for idx, task := range list.Items {
		listResponse[idx] = controller.Service.PrepareTaskResponse(&task)
	}

	ctx.JSON(200, response.GetListResponse{
		Count: int(list.Count),
		Items: listResponse[:],
	})
}

func (controller *TaskControllerImpl) Create(ctx *gin.Context) {
	reqBody := request.CreateTaskRequest{}
	jsonError := ctx.ShouldBindJSON(&reqBody)

	if jsonError != nil {
		ctx.JSON(400, response.ApiErrorResponse{
			Error: "Failed to parse JSON",
		})
		return
	}

	validationError := controller.Validate.Struct(reqBody)

	if validationError != nil {
		ctx.JSON(400, response.ApiErrorResponse{
			Error: "Not valid JSON",
		})
		return
	}

	newTask := models.Task{}
	newTask.Name = reqBody.Name

	err := controller.Service.Create(&newTask)
	if err != nil {
		ctx.JSON(400, response.ApiErrorResponse{
			Error: err.Error(),
		})
		return
	}

	ctx.Status(204)
}

func (controller *TaskControllerImpl) Update(ctx *gin.Context) {
	reqBody := request.UpdateTaskRequest{}
	jsonError := ctx.ShouldBindJSON(&reqBody)
	idParam, idParamExist := ctx.Params.Get("id")
	var id int

	if !idParamExist {
		ctx.JSON(400, response.ApiErrorResponse{
			Error: "Missed param :id",
		})
		return
	}

	id, parseErr := strconv.Atoi(idParam)

	if parseErr != nil {
		ctx.JSON(400, response.ApiErrorResponse{
			Error: "Unexpected value in param :id",
		})
		return
	}

	if jsonError != nil {
		ctx.JSON(400, response.ApiErrorResponse{
			Error: "Failed to parse JSON",
		})
		return
	}

	validationError := controller.Validate.Struct(reqBody)

	if validationError != nil {
		ctx.JSON(400, response.ApiErrorResponse{
			Error: "Not valid JSON",
		})
		return
	}

	updateTask := models.Task{}

	updateTask.Status = reqBody.Status
	updateTask.Name = reqBody.Name
	updateTask.ID = uint(id)

	err := controller.Service.Update(&updateTask)
	if err != nil {
		ctx.JSON(400, response.ApiErrorResponse{
			Error: err.Error(),
		})
		return
	}

	ctx.Status(204)
}

func (controller *TaskControllerImpl) Delete(ctx *gin.Context) {
	idParam, available := ctx.Params.Get("id")
	if !available {
		ctx.JSON(400, response.ApiErrorResponse{
			Error: "param :id is not provided",
		})
		return
	}
	id, parseFail := strconv.Atoi(idParam)
	if parseFail != nil {
		ctx.JSON(400, response.ApiErrorResponse{
			Error: "Unexpected value in param :id",
		})
		return
	}

	controller.Service.Delete(uint(id))
	ctx.Status(204)
}

func NewTaskController(service services.TaskService, validate *validator.Validate) TaskController {
	return &TaskControllerImpl{
		Service:  service,
		Validate: validate,
	}
}
