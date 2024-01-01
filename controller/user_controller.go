package controller

import (
	usersRequest "go-project/data/requests/users"
	response "go-project/data/responses"
	"go-project/helper"
	"go-project/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	UserService service.UserService
	Validate    *validator.Validate
}

func UserControllerExecution() *UserController {
	service := service.UserServiceImplExecution()
	validate := validator.New()
	return &UserController{
		UserService: service,
		Validate:    validate,
	}
}

func (controller *UserController) Create(ctx *gin.Context) {
	createUserRequest := usersRequest.CreateUserRequest{}
	err := ctx.ShouldBindJSON(&createUserRequest)
	if err != nil {
		helper.Error(ctx, err)
		return
	}
	err = controller.Validate.Struct(createUserRequest)
	if err != nil {
		helper.Error(ctx, err)
		return
	}

	user, err := controller.UserService.Create(createUserRequest)
	if err != nil {
		helper.Error(ctx, err)
	}
	data := helper.RespondWithJSON(user)

	ctx.JSON(http.StatusOK, data)
}

func (controller *UserController) Update(ctx *gin.Context) {
	UpdateUserRequest := usersRequest.UpdateUserRequest{}
	err := ctx.ShouldBindJSON(&UpdateUserRequest)
	if err != nil {
		helper.Error(ctx, err)
		return
	}
	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		helper.Error(ctx, err)
		return
	}
	UpdateUserRequest.Id = id

	user := controller.UserService.Update(UpdateUserRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   user,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *UserController) Delete(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	controller.UserService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *UserController) FindById(ctx *gin.Context) {
	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)

	userResponse := controller.UserService.FindById(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   userResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *UserController) FindAll(ctx *gin.Context) {
	userResponse := controller.UserService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   userResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
