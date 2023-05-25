package store

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int
	Username  string
	Name      string
	Bio       sql.NullString
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
	ID       int
	Login    string
	Username string
}

type UserStore interface {
	CreateUser(opts CreateUserOpts) (User, error)
	GetUserByLogin(opts GetUserOpts) (User, error)
	GetUserByID(opts GetUserOpts) (User, error)
	GetUserByUsername(opts GetUserOpts) (User, error)
}
