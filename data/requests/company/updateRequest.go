package companyRequest

type UpdateCompanyRequest struct {
	CompanyName string `validate:"omitempty,min=1,max=200" json:"company_name"`
	Email       string `validate:"omitempty" json:"email"`
	DirectorFIO string `validate:"omitempty" json:"director_fio"`
	Address     string `validate:"omitempty" json:"address"`
	Website     string `validate:"required"  json:"website"`
}



