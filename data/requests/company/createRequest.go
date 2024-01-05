package companyRequest

type CreateCompanyRequest struct {
	CompanyName string `validate:"required,min=1,max=200" json:"company_name"`
	UserID      int    `validate:"required" json:"user_id"`
	Email       string `json:"email" validate:"required,email"`
	DirectorFIO string `validate:"required" json:"director_fio"`
	Address     string `validate:"required" json:"address"`
	Website     string `validate:"required" json:"website"`
}
