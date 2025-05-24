package models


type User struct {
	ID		uint `json:"id"`
	Name  string `json:"username"`
	Email string `json:"email"`
	PassHash string `json:"-"`
}
