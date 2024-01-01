package service

import (
	usersRequest "go-project/data/requests/users"
	usersResponse "go-project/data/responses/users"
	"go-project/helper"
	"go-project/infrastructure/interfaces"
	"go-project/model"
	"go-project/repository"
	"go-project/utils"
)

type UserService interface {
	Create(user usersRequest.CreateUserRequest) (userModel model.User, err error)
	Update(user usersRequest.UpdateUserRequest) model.User
	Delete(userId int)
	FindById(userId int) usersResponse.BaseResponse
	FindAll() []usersResponse.BaseResponse
}

type UserServiceImpl struct {
	UserRepository interfaces.UserRepository
}

func UserServiceImplExecution() UserService {
	userRepository := repository.UserREpositoryImplExecution()
	return &UserServiceImpl{
		UserRepository: userRepository,
	}
}

func (t *UserServiceImpl) Create(user usersRequest.CreateUserRequest) (users model.User, err error) {
	password, err := utils.HashPassword(user.Password)
	if err != nil {
		return users, err
	}
	userModel := model.User{
		Name:     user.Name,
		Email:    user.Email,
		Role:     user.Role,
		Password: password,
	}
	t.UserRepository.Save(userModel)
	return userModel, nil
}

func (t *UserServiceImpl) Delete(userId int) {
	t.UserRepository.Delete(userId)
}

func (t *UserServiceImpl) FindAll() []usersResponse.BaseResponse {
	result := t.UserRepository.FindAll()

	var users []usersResponse.BaseResponse
	for _, value := range result {
		user := usersResponse.BaseResponse{
			Id:    value.ID,
			Name:  value.Name,
			Email: value.Email,
			Role:  value.Role,
		}
		users = append(users, user)
	}

	return users
}

func (t *UserServiceImpl) FindById(userId int) usersResponse.BaseResponse {
	userData, err := t.UserRepository.FindById(userId)
	helper.ErrorPanic(err)

	userResponse := usersResponse.BaseResponse{
		Id:    userData.ID,
		Name:  userData.Name,
		Email: userData.Email,
		Role:  userData.Role,
	}
	return userResponse
}

func (t *UserServiceImpl) Update(users usersRequest.UpdateUserRequest) model.User {
	userData, err := t.UserRepository.FindById(users.Id)
	helper.ErrorPanic(err)

	if users.Name != "" {
		userData.Name = users.Name
	}
	if users.Password != "" {
		userData.Password = users.Password
	}
	if users.Email != "" {
		userData.Email = users.Email
	}
	if users.Role != "" {
		userData.Role = users.Role
	}

	user := t.UserRepository.Update(userData)

	return user
}
