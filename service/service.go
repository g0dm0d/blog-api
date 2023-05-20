package service

import (
	"blog-api/service/user"
	"blog-api/store"
	"blog-api/tools/tokenmanager"
)

type Service struct {
	User user.User
}

func New(userStore store.UserStore, sessionStore store.SessionStore, token tokenmanager.Tool) *Service {
	return &Service{
		User: user.New(userStore, sessionStore, token),
	}
}
