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

func (s *SessionStore) UpdateSession(opts store.UpdateSessionOpts) (int, error) {
	var UserID int
	result := s.db.QueryRow("SELECT * FROM update_session($1, $2)", opts.NewToken, opts.OldToken)
	err := result.Scan(&UserID)
	if err != nil {
		return 0, err
	}
	return UserID, err
}

func (s *SessionStore) ClearExpSession() error {
	_, err := s.db.Exec("CALL delete_expired_sessions()")
	return err
}
