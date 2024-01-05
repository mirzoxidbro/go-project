package CompanyResponse

type CompanyResponse struct {
	Id          int    `json:"id"`
	UserID      int    `json:"user_id"`
	CompanyName string `json:"company_name"`
	DirectorFIO string `json:"director_fio"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Website     string `json:"website"`
}
