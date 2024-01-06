package controller

import (
	EmployeeRequest "go-project/data/requests/employee"
	"go-project/helper"
	"go-project/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type EmployeeController struct {
	EmployeeService service.EmployeeService
	Validate        *validator.Validate
}

func EmployeeControllerExecution() *EmployeeController {
	service := service.EmployeeServiceImplExecution()
	validate := validator.New()
	return &EmployeeController{
		EmployeeService: service,
		Validate:        validate,
	}
}

func (controller *EmployeeController) FindAll(ctx *gin.Context) {
	employeeResponse := controller.EmployeeService.FindAll()

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, helper.RespondWithJSON(employeeResponse))
}

func (controller *EmployeeController) Create(ctx *gin.Context) {
	createEmployeeRequest := EmployeeRequest.CreateEmployeeRequest{}
	err := ctx.ShouldBindJSON(&createEmployeeRequest)
	if err != nil {
		helper.Error(ctx, err)
		return
	}
	err = controller.Validate.Struct(createEmployeeRequest)
	if err != nil {
		helper.Error(ctx, err)
		return
	}

	company, err := controller.EmployeeService.Create(createEmployeeRequest)
	if err != nil {
		helper.Error(ctx, err)
	}
	data := helper.RespondWithJSON(company)

	ctx.JSON(http.StatusOK, data)
}

func (controller *EmployeeController) FindById(ctx *gin.Context) {
	employeeId := ctx.Param("employeeId")
	id, err := strconv.Atoi(employeeId)
	helper.ErrorPanic(err)

	employeeResponse := controller.EmployeeService.FindById(id)
	response := helper.RespondWithJSON(employeeResponse)

	ctx.JSON(http.StatusOK, response)
}

func (controller *EmployeeController) Update(ctx *gin.Context) {
	UpdateEmployeeRequest := EmployeeRequest.UpdateEmployeeRequest{}
	err := ctx.ShouldBindJSON(&UpdateEmployeeRequest)
	if err != nil {
		helper.Error(ctx, err)
		return
	}
	employeeId := ctx.Param("employeeId")
	id, err := strconv.Atoi(employeeId)
	if err != nil {
		helper.Error(ctx, err)
		return
	}

	user := controller.EmployeeService.Update(id, UpdateEmployeeRequest)

	response := helper.RespondWithJSON(user)

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, response)
}

func (controller *EmployeeController) Delete(ctx *gin.Context) {
	employeeId := ctx.Param("employeeId")
	id, err := strconv.Atoi(employeeId)
	helper.ErrorPanic(err)
	controller.EmployeeService.Delete(id)

	response := helper.RespondWithJSON(nil)

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, response)
}
