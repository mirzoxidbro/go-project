package controller

import (
	companyRequest "go-project/data/requests/company"
	response "go-project/data/responses"
	"go-project/helper"
	"go-project/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CompanyController struct {
	CompanyService service.CompanyService
	Validate       *validator.Validate
}

func CompanyControllerExecution() *CompanyController {
	service := service.CompanyServiceImplExecution()
	validate := validator.New()
	return &CompanyController{
		CompanyService: service,
		Validate:       validate,
	}
}

func (controller *CompanyController) Create(ctx *gin.Context) {
	createCompanyRequest := companyRequest.CreateCompanyRequest{}
	err := ctx.ShouldBindJSON(&createCompanyRequest)
	if err != nil {
		helper.Error(ctx, err)
		return
	}
	err = controller.Validate.Struct(createCompanyRequest)
	if err != nil {
		helper.Error(ctx, err)
		return
	}

	company, err := controller.CompanyService.Create(createCompanyRequest)
	if err != nil {
		helper.Error(ctx, err)
	}
	data := helper.RespondWithJSON(company)

	ctx.JSON(http.StatusOK, data)
}

func (controller *CompanyController) Update(ctx *gin.Context) {
	UpdateCompanyRequest := companyRequest.UpdateCompanyRequest{}
	err := ctx.ShouldBindJSON(&UpdateCompanyRequest)
	if err != nil {
		helper.Error(ctx, err)
		return
	}
	companyId := ctx.Param("companyId")
	id, err := strconv.Atoi(companyId)
	if err != nil {
		helper.Error(ctx, err)
		return
	}

	user := controller.CompanyService.Update(id, UpdateCompanyRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   user,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *CompanyController) Delete(ctx *gin.Context) {
	companyId := ctx.Param("companyId")
	id, err := strconv.Atoi(companyId)
	helper.ErrorPanic(err)
	controller.CompanyService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *CompanyController) FindById(ctx *gin.Context) {
	companyId := ctx.Param("companyId")
	id, err := strconv.Atoi(companyId)
	helper.ErrorPanic(err)

	userResponse := controller.CompanyService.FindById(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   userResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *CompanyController) FindAll(ctx *gin.Context) {
	userResponse := controller.CompanyService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   userResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
