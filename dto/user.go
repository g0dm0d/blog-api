package dto

import "blog-api/model"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     int16  `json:"role"`
}

type UserPublic struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Bio      string `json:"bio"`
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
		ID:       u.ID,
		Username: u.Username,
		Name:     u.Name,
		Bio:      u.Bio,
	}
}

func NewUsersPublic(u []model.User) []UserPublic {
	var usersPublic []UserPublic
	for i := range u {
		usersPublic = append(usersPublic, NewUserPublic(u[i]))
	}
	return usersPublic
}
