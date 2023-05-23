package store

import "time"

type Session struct {
	Id           string
	UserId       string
	RefreshToken string
	ExpiresAt    time.Time
	CreatedAt    time.Time
}

type CreateSessionOpts struct {
	UserId       int
	RefreshToken string
}

type UpdateSessionOpts struct {
	NewToken string
	OldToken string
}

type SessionStore interface {
	CreateSession(opts CreateSessionOpts) error
	UpdateSession(opts UpdateSessionOpts) (int, error)
	ClearExpSession() error
}
