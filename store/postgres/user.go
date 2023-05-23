package postgres

import (
	"blog-api/store"
	"database/sql"
)

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) store.UserStore {
	return &UserStore{
		db: db,
	}
}

func (s *UserStore) CreateUser(opts store.CreateUserOpts) (store.User, error) {
	_, err := s.db.Exec("CALL create_user($1, $2, $3)",
		opts.Username, opts.Email, opts.Password)

	return store.User{}, err
}

// GetUserByLogin Login is email or username
func (s *UserStore) GetUserByLogin(opts store.GetUserOpts) (store.User, error) {
	var user store.User
	req := s.db.QueryRow("SELECT * FROM get_user_by_email_or_username($1)", opts.Login)
	err := req.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role, &user.CreatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *UserStore) GetUserByID(userID int) (store.User, error) {
	var user store.User
	req := s.db.QueryRow("SELECT * FROM get_user_by_id($1)", userID)
	err := req.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role, &user.CreatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}
