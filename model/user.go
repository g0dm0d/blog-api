package model

import (
	"blog-api/store"
	"time"
)

type User struct {
	ID        int
	Username  string
	Name      string
	Bio       string
	Email     string
	Password  string
	Role      int16
	CreatedAt time.Time
}

func NewUser(u store.User) User {
	return User{
		ID:        u.ID,
		Username:  u.Username,
		Name:      u.Name,
		Bio:       u.Bio.String,
		Email:     u.Email,
		Password:  u.Password,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
	}
}

func NewUsers(u []store.User) []User {
	var users []User
	for i := range u {
		users = append(users, NewUser(u[i]))
	}
	return users
}
