package response

import "api/internal/models"

type UserResponse struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
}

func NewUserResponse(u *models.User) UserResponse {
	return UserResponse{
		Name:      u.Name,
		Email:     u.Email,
		Telephone: u.Telephone,
	}
}
