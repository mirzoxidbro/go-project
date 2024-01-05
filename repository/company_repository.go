package repository

import (
	"errors"
	"go-project/helper"
	"go-project/infrastructure/interfaces"
	"go-project/initializers"
	"go-project/model"

	"gorm.io/gorm"
)

type CompanyRepositoryImpl struct {
	Db *gorm.DB
}

func CompanyRepositoryImplExecution() interfaces.CompanyRepository {
	db := initializers.DB
	return &CompanyRepositoryImpl{Db: db}
}

func (c *CompanyRepositoryImpl) Create(company model.Company) (Company model.Company, err error) {
	result := c.Db.Create(&company)
	if result.Error != nil {
		helper.ErrorPanic(result.Error)
		return company, result.Error
	}
	return company, nil
}

func (c *CompanyRepositoryImpl) FindAll() []model.Company {
	var company []model.Company
	result := c.Db.Find(&company)
	helper.ErrorPanic(result.Error)
	return company
}

func (c *CompanyRepositoryImpl) FindById(companyId int) (Company model.Company, err error) {
	var company model.Company
	result := c.Db.Find(&company, companyId)
	if result != nil {
		return company, nil
	} else {
		return company, errors.New("Company is not found")
	}
}

func (c *CompanyRepositoryImpl) Update(company model.Company) model.Company {
	updateCompany := map[string]interface{}{}

	if company.CompanyName != "" {
		updateCompany["CompanyName"] = company.CompanyName
	}
	if company.Website != "" {
		updateCompany["Website"] = company.Website
	}
	if company.Email != "" {
		updateCompany["Email"] = company.Email
	}
	if company.DirectorFIO != "" {
		updateCompany["DirectorFIO"] = company.DirectorFIO
	}
	if company.Address != "" {
		updateCompany["Address"] = company.Address
	}

	c.Db.Model(&company).Updates(updateCompany)

	c.Db.First(&company, company.ID)
	return company
}

func (c *CompanyRepositoryImpl) Delete(companyId int) {
	var company model.Company
	result := c.Db.Where("id = ?", companyId).Delete(&company)
	helper.ErrorPanic(result.Error)
}
