package request

type LoginRequest struct {
	Username string `json:"username" validate:"omitempty"`
	Email    string `json:"email" validate:"omitempty,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Username *string `json:"username" validate:"omitempty"`
	Password string  `json:"password" validate:"required"`
	Email    *string `json:"email" validate:"omitempty,email"`
	Name     string  `json:"name" validate:"required,min=2,max=255"`
}
