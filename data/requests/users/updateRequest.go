package usersRequest

type UpdateUserRequest struct {
	Name     string `validate:"omitempty,min=1,max=200" json:"name"`
	Email    string `validate:"omitempty" json:"email"`
	Password string `validate:"omitempty" json:"password"`
	Role     string `validate:"omitempty" json:"role"`
	Id       int    `validate:"required"`
}
