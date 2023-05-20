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
	UserId       string
	RefreshToken string
}

type SessionStore interface {
	CreateSession(opts CreateSessionOpts) error
	ClearExpSession() error
}
