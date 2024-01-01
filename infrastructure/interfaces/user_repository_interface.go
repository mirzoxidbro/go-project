package interfaces

import "go-project/model"

type UserRepository interface {
	Save(user model.User) (users model.User, err error)
	Update(user model.User) model.User
	Delete(userId int)
	FindById(userId int) (user model.User, err error)
	FindAll() []model.User
}
