package request

type LoginRequest struct {
	Username *string `json:"username" validate:"omitempty"`
	Email    *string `json:"email" validate:"omitempty,email"`
	Password string  `json:"password" validate:"required"`
}
