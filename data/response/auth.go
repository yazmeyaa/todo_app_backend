package response

import "github.com/yazmeyaa/todo_app_backend/models"

type userRespones struct {
	ID       int     `json:"id"`
	Username *string `json:"usernmae"`
	Email    *string `json:"email"`
	Name     string  `json:"name"`
}

type LoginResponse struct {
	User  userRespones `json:"user"`
	Token string       `json:"token"`
}

func NewLoginResponse(user *models.User, token string) LoginResponse {
	return LoginResponse{
		User: userRespones{
			ID:       int(user.ID),
			Username: user.Username,
			Email:    user.Email,
			Name:     user.Name,
		},
		Token: token,
	}
}
