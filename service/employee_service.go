package service

import (
	EmployeeRequest "go-project/data/requests/employee"
	EmployeeResponse "go-project/data/responses/employee"
	"go-project/helper"
	"go-project/infrastructure/interfaces"
	"go-project/model"
	"go-project/repository"
)

type EmployeeService interface {
	Create(employee EmployeeRequest.CreateEmployeeRequest) (Employee *model.Employee, err error)
	Update(employeeId int, company EmployeeRequest.UpdateEmployeeRequest) model.Employee
	Delete(companyId int)
	FindById(companyId int) EmployeeResponse.EmployeeResponse
	FindAll() []EmployeeResponse.EmployeeResponse
}

type EmployeeServiceImpl struct {
	EmployeeRepository interfaces.EmployeeRepository
}

func EmployeeServiceImplExecution() EmployeeService {
	employeeRepository := repository.EmployeeRepositoryImplExecution()
	return &EmployeeServiceImpl{
		EmployeeRepository: employeeRepository,
	}
}

func (e *EmployeeServiceImpl) FindAll() []EmployeeResponse.EmployeeResponse {
	result := e.EmployeeRepository.FindAll()

	var employees []EmployeeResponse.EmployeeResponse
	for _, value := range result {
		employee := EmployeeResponse.EmployeeResponse{
			ID:             value.ID,
			CompanyID:      value.CompanyID,
			Passport:       value.Passport,
			Surname:        value.Surname,
			Firstname:      value.Firstname,
			PatronymicName: value.PatronymicName,
			Position:       value.Position,
			PhoneNumber:    value.PhoneNumber,
			Address:        value.Address,
			CreatedAt:      value.CreatedAt,
			UpdaredAt:      value.UpdatedAt,
		}
		employees = append(employees, employee)
	}

	return employees
}

func (e *EmployeeServiceImpl) Create(employee EmployeeRequest.CreateEmployeeRequest) (Employee *model.Employee, err error) {

	employeeModel := model.Employee{
		CompanyID:      employee.CompanyID,
		Passport:       employee.Passport,
		Surname:        employee.Surname,
		Firstname:      employee.Firstname,
		PatronymicName: employee.PatronymicName,
		Position:       employee.Position,
		PhoneNumber:    employee.PhoneNumber,
		Address:        employee.Address,
	}
	e.EmployeeRepository.Create(employeeModel)
	return &employeeModel, nil
}

func (e *EmployeeServiceImpl) FindById(employeeId int) EmployeeResponse.EmployeeResponse {
	result, err := e.EmployeeRepository.FindById(employeeId)
	helper.ErrorPanic(err)

	employee := EmployeeResponse.EmployeeResponse{
		ID:             result.ID,
		CompanyID:      result.CompanyID,
		Passport:       result.Passport,
		Surname:        result.Surname,
		Firstname:      result.Firstname,
		PatronymicName: result.PatronymicName,
		Position:       result.Position,
		PhoneNumber:    result.PhoneNumber,
		Address:        result.Address,
		CreatedAt:      result.CreatedAt,
		UpdaredAt:      result.UpdatedAt,
	}
	return employee
}

func (e *EmployeeServiceImpl) Update(employeeId int, employee EmployeeRequest.UpdateEmployeeRequest) model.Employee {
	employeeData, err := e.EmployeeRepository.FindById(employeeId)
	helper.ErrorPanic(err)

	if employeeData.CompanyID != 0 {
		employeeData.CompanyID = employee.CompanyID
	}
	if employeeData.Passport != "" {
		employeeData.Passport = employee.Passport
	}
	if employeeData.Surname != "" {
		employeeData.Surname = employee.Surname
	}
	if employeeData.Firstname != "" {
		employeeData.Firstname = employee.Firstname
	}
	if employeeData.PatronymicName != "" {
		employeeData.Address = employee.PatronymicName
	}
	if employeeData.Position != "" {
		employeeData.Position = employee.Position
	}
	if employeeData.PhoneNumber != "" {
		employeeData.PhoneNumber = employee.PhoneNumber
	}
	if employeeData.Address != "" {
		employeeData.Address = employee.Address
	}

	updatedData := e.EmployeeRepository.Update(employeeData)

	return updatedData
}

func (e *EmployeeServiceImpl) Delete(employeeId int) {
	e.EmployeeRepository.Delete(employeeId)
}
