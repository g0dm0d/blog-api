package store

import (
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

type CreateUserOpts struct {
	Username string
	Email    string
	Password string
}

type GetUserOpts struct {
	Login string
}

type UserStore interface {
	CreateUser(opts CreateUserOpts) (User, error)
	GetUserByLogin(opts GetUserOpts) (User, error)
	GetUserByID(userID int) (User, error)
}
