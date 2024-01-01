package usersRequest

type CreateUserRequest struct {
	Name     string `validate:"required,min=1,max=200" json:"name"`
	Email    string `json:"email" validate:"required,email"`
	Password string `validate:"required" json:"password"`
	Role     string `json:"role" validate:"required,oneof=admin company"`
}
