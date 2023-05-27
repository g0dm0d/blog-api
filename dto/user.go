package dto

import (
	"blog-api/model"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	Avatar    string    `json:"avatar"`
	Bio       string    `json:"bio"`
	Email     string    `json:"email"`
	Role      int16     `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

type UserPublic struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	Avatar    string    `json:"avatar"`
	Bio       string    `json:"bio"`
	CreatedAt time.Time `json:"created_at"`
}

func NewUser(u model.User) User {
	return User{
		ID:        u.ID,
		Username:  u.Username,
		Name:      u.Name,
		Avatar:    u.Avatar,
		Bio:       u.Bio,
		Email:     u.Email,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
	}
}

func NewUserPublic(u model.User) UserPublic {
	return UserPublic{
		ID:        u.ID,
		Username:  u.Username,
		Name:      u.Name,
		Avatar:    u.Avatar,
		Bio:       u.Bio,
		CreatedAt: u.CreatedAt,
	}
}

func NewUsersPublic(u []model.User) []UserPublic {
	var usersPublic []UserPublic
	for i := range u {
		usersPublic = append(usersPublic, NewUserPublic(u[i]))
	}
	return usersPublic
}
