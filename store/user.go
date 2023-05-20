package store

import (
	"time"
)

type User struct {
	Id        string
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
	GetUser(opts GetUserOpts) (User, error)
}
