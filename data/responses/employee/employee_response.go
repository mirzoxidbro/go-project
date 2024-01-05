package EmployeeResponse

import "time"

type EmployeeResponse struct {
	ID             int       `json:"id"`
	CompanyID      int       `json:"company_id"`
	Passport       string    `json:"passport"`
	Surname        string    `json:"surname"`
	Firstname      string    `json:"firstname"`
	PatronymicName string    `json:"patronymic_name"`
	Position       string    `json:"position"`
	PhoneNumber    string    `json:"phone_number"`
	Address        string    `json:"address"`
	CreatedAt      time.Time `json:"created_at"`
	UpdaredAt      time.Time `json:"updated_at"`
}
