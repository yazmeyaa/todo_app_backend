package request

type CreateTaskRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdateTaskRequest struct {
	Status string `json:"status" validate:"oneof=created not_done in_work done"`
	Name   string `json:"name" validate:"min=1,max=255"`
}
