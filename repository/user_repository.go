package repository

import (
	"errors"
	"go-project/helper"
	"go-project/infrastructure/interfaces"
	"go-project/initializers"
	"go-project/model"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func UserRepositoryImplExecution() interfaces.UserRepository {
	db := initializers.DB
	return &UserRepositoryImpl{Db: db}
}

func (t *UserRepositoryImpl) Save(users model.User) (user model.User, err error) {
	result := t.Db.Create(&users)
	if result.Error != nil {
		helper.ErrorPanic(result.Error)
		return user, result.Error
	}
	return users, nil
}

func (t *UserRepositoryImpl) FindAll() []model.User {
	var user []model.User
	result := t.Db.Find(&user)
	helper.ErrorPanic(result.Error)
	return user
}

func (t *UserRepositoryImpl) FindById(userId int) (users model.User, err error) {
	var user model.User
	result := t.Db.Find(&user, userId)
	if result != nil {
		return user, nil
	} else {
		return user, errors.New("tag is not found")
	}
}

func (t *UserRepositoryImpl) Update(user model.User) model.User {
	updateUser := map[string]interface{}{
		"Name":     user.Name,
		"Role":     user.Role,
		"Email":    user.Email,
		"Password": user.Password,
	}
	t.Db.Model(&user).Updates(updateUser)

	var updatedUser model.User
	t.Db.First(&updatedUser, user.ID)
	return updatedUser
}

func (t *UserRepositoryImpl) Delete(userId int) {
	var user model.User
	result := t.Db.Where("id = ?", userId).Delete(&user)
	helper.ErrorPanic(result.Error)
}
