package entity

import "time"

type Applicants struct {
	ID          int        `json:"id"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	Email       string     `json:"email"`
	Phone       string     `json:"phone"`
	HomeAddress string     `json:"home_address"`
	Title       string     `json:"title"`
	YearsOfExp  int        `json:"years_of_exp"`
	CreatedAt   time.Time  `json:"created_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}
