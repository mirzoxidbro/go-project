package repository

import (
	"errors"
	"go-project/helper"
	"go-project/infrastructure/interfaces"
	"go-project/initializers"
	"go-project/model"

	"gorm.io/gorm"
)

type EmployeeRepositoryImpl struct {
	Db *gorm.DB
}

func EmployeeRepositoryImplExecution() interfaces.EmployeeRepository {
	db := initializers.DB
	return &EmployeeRepositoryImpl{Db: db}
}

func (e *EmployeeRepositoryImpl) FindAll() []model.Employee {
	var employee []model.Employee
	result := e.Db.Find(&employee)
	helper.ErrorPanic(result.Error)
	return employee
}

func (e *EmployeeRepositoryImpl) Create(employee model.Employee) (Employee model.Employee, err error) {
	result := e.Db.Create(&employee)
	if result.Error != nil {
		helper.ErrorPanic(result.Error)
		return employee, result.Error
	}
	return employee, nil
}

func (e *EmployeeRepositoryImpl) FindById(employeeId int) (Employee model.Employee, err error) {
	var employee model.Employee
	result := e.Db.Find(&employee, employeeId)
	if result != nil {
		return employee, nil
	} else {
		return employee, errors.New("Company is not found")
	}
}

func (e *EmployeeRepositoryImpl) Update(employee model.Employee) (Employee model.Employee) {
	updateEmployee := map[string]interface{}{}
	if employee.Passport != "" {
		updateEmployee["Passport"] = employee.Passport
	}
	if employee.Surname != "" {
		updateEmployee["Surname"] = employee.Surname
	}
	if employee.Firstname != "" {
		updateEmployee["Firstname"] = employee.Firstname
	}
	if employee.PatronymicName != "" {
		updateEmployee["PatronymicName"] = employee.PatronymicName
	}
	if employee.Position != "" {
		updateEmployee["Position"] = employee.Position
	}

	e.Db.Model(&employee).Updates(updateEmployee)

	e.Db.First(&employee, employee.ID)
	return employee
}

func (e *EmployeeRepositoryImpl) Delete(employeeId int) {
	var employee model.Employee

	result := e.Db.Where("id = ?", employeeId).Delete(employee)
	helper.ErrorPanic(result.Error)
}
