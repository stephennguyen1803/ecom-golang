package model

type User struct {
	// User model
}

func NewUser() *User {
	return &User{}
}

func (u *User) GetUserName() string {
	return "Anh Dung Nguyen"
}
