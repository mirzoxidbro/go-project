package EmployeeRequest

type CreateEmployeeRequest struct {
	CompanyID      int    `validate:"required" json:"company_id"`
	Passport       string `validate:"required" json:"passport"`
	Surname        string `validate:"required" json:"surname"`
	Firstname      string `validate:"required" json:"firstname"`
	PatronymicName string `validate:"required" json:"patronymic_name"`
	Position       string `validate:"required" json:"position"`
	PhoneNumber    string `validate:"required" json:"phone_number"`
	Address        string `validate:"required" json:"address"`
}
