package dto

import "blog-api/model"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     int16  `json:"role"`
}

type UserPublic struct {
	Username string `json:"username"`
}

func NewUser(u model.User) User {
	return User{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Role:     u.Role,
	}
}

func NewUserPublic(u model.User) UserPublic {
	return UserPublic{
		Username: u.Username,
	}
}
