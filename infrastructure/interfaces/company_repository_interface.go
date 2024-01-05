package interfaces

import "go-project/model"

type CompanyRepository interface {
	Create(Company model.Company) (company model.Company, err error)
	Update(company model.Company) model.Company
	Delete(companyId int)
	FindById(companyId int) (company model.Company, err error)
	FindAll() []model.Company
}
