package request

type CreateUserRequest struct {
	Username string `json:"username" validate:"omitempty,min=8,max=32"`
	Email    string `json:"email" validate:"omitempty,email"`
	Name     string `json:"name" validate:"required,max=100"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}
