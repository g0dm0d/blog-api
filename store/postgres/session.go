package postgres

import (
	"blog-api/store"
	"database/sql"
)

type SessionStore struct {
	db *sql.DB
}

func NewSessionStore(db *sql.DB) store.SessionStore {
	return &SessionStore{
		db: db,
	}
}

func (s *SessionStore) CreateSession(opts store.CreateSessionOpts) error {
	_, err := s.db.Exec("CALL create_session($1, $2)", opts.UserId, opts.RefreshToken)
	return err
}

func (s *SessionStore) ClearExpSession() error {
	_, err := s.db.Exec("CALL delete_expired_sessions()")
	return err
}
