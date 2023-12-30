package usersRequest

type CreateUserRequest struct {
	Name     string `validate:"required,min=1,max=200" json:"name"`
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
	Role     string `validate:"required" json:"role"`
}
