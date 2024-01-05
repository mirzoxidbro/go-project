package EmployeeRequest

type UpdateEmployeeRequest struct {
	CompanyID      int    `validate:"omitempty" json:"company_id"`
	Passport       string `validate:"omitempty" json:"passport"`
	Surname        string `validate:"omitempty" json:"surname"`
	Firstname      string `validate:"omitempty" json:"firstname"`
	PatronymicName string `validate:"omitempty" json:"patronymic_name"`
	Position       string `validate:"omitempty" json:"position"`
	PhoneNumber    string `validate:"omitempty" json:"phone_number"`
	Address        string `validate:"omitempty" json:"address"`
}
