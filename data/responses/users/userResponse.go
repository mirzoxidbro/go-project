package usersResponse

import "go-project/model"

type UserResponse struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Role    string `json:"role"`
	Company model.Company
}
