package presentation

type ApplicantRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	HomeAddress string `json:"home_address"`
	Title       string `json:"title"`
	YearsOfExp  int    `json:"years_of_exp"`
}
