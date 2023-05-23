package model

import (
	"blog-api/store"
	"time"
)

type User struct {
	ID        int
	Username  string
	Email     string
	Password  string
	Role      int16
	CreatedAt time.Time
}

func NewUser(u store.User) User {
	return User{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
	}
}
