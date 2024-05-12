package response

import "time"

type TaskResponse struct {
	ID        int        `json:"id"`
	DeletedAt *time.Time `json:"deletedAt"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	Name      string     `json:"name"`
	Status    string     `json:"status"`
}

type CreateTaskResponse struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"createdAt"`
}

type GetByIdResponse struct {
	Task TaskResponse `json:"task"`
}

type GetListResponse struct {
	Count int            `json:"count"`
	Items []TaskResponse `json:"items"`
}
