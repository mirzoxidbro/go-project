package service

import (
	companyRequest "go-project/data/requests/company"
	CompanyResponse "go-project/data/responses/company"
	"go-project/helper"
	"go-project/infrastructure/interfaces"
	"go-project/model"
	"go-project/repository"
)

type CompanyService interface {
	Create(company companyRequest.CreateCompanyRequest) (companyModel *model.Company, err error)
	Update(companyId int, company companyRequest.UpdateCompanyRequest) model.Company
	Delete(companyId int)
	FindById(companyId int) CompanyResponse.CompanyResponse
	FindAll() []CompanyResponse.CompanyResponse
}

type CompanyServiceImpl struct {
	CompanyRepository interfaces.CompanyRepository
}

func CompanyServiceImplExecution() CompanyService {
	companyRepository := repository.CompanyRepositoryImplExecution()
	return &CompanyServiceImpl{
		CompanyRepository: companyRepository,
	}
}

func (c *CompanyServiceImpl) FindAll() []CompanyResponse.CompanyResponse {
	result := c.CompanyRepository.FindAll()

	var companies []CompanyResponse.CompanyResponse
	for _, value := range result {
		company := CompanyResponse.CompanyResponse{
			Id:          value.ID,
			UserID:      value.UserID,
			CompanyName: value.CompanyName,
			Email:       value.Email,
			DirectorFIO: value.DirectorFIO,
			Address:     value.Address,
			Website:     value.Website,
		}
		companies = append(companies, company)
	}

	return companies
}

func (c *CompanyServiceImpl) Create(company companyRequest.CreateCompanyRequest) (Company *model.Company, err error) {

	companyModel := model.Company{
		CompanyName: company.CompanyName,
		UserID:      company.UserID,
		Email:       company.Email,
		DirectorFIO: company.DirectorFIO,
		Address:     company.Address,
		Website:     company.Website,
	}
	c.CompanyRepository.Create(companyModel)
	return &companyModel, nil
}

func (c *CompanyServiceImpl) FindById(companyId int) CompanyResponse.CompanyResponse {
	result, err := c.CompanyRepository.FindById(companyId)
	helper.ErrorPanic(err)

	company := CompanyResponse.CompanyResponse{
		Id:          result.ID,
		CompanyName: result.CompanyName,
		Email:       result.Email,
		DirectorFIO: result.DirectorFIO,
		Address:     result.Address,
		Website:     result.Website,
		UserID:      result.UserID,
	}
	return company
}

func (c *CompanyServiceImpl) Update(companyId int, company companyRequest.UpdateCompanyRequest) model.Company {
	companyData, err := c.CompanyRepository.FindById(companyId)
	helper.ErrorPanic(err)

	if companyData.CompanyName != "" {
		companyData.CompanyName = company.CompanyName
	}
	if companyData.Website != "" {
		companyData.Website = company.Website
	}
	if companyData.Email != "" {
		companyData.Email = company.Email
	}
	if companyData.DirectorFIO != "" {
		companyData.DirectorFIO = company.DirectorFIO
	}
	if companyData.Address != "" {
		companyData.Address = company.Address
	}

	updatedData := c.CompanyRepository.Update(companyData)

	return updatedData
}

func (c *CompanyServiceImpl) Delete(companyId int) {
	c.CompanyRepository.Delete(companyId)
}
