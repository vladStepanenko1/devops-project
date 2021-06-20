package repository

type User struct {
	Id          int    `json:"-"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}
