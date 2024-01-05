package interfaces

import "go-project/model"

type EmployeeRepository interface {
	Create(employee model.Employee) (Employee model.Employee, err error)
	Update(employee model.Employee) model.Employee
	Delete(employeeId int)
	FindById(employeeId int) (employee model.Employee, err error)
	FindAll() []model.Employee
}
