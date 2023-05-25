package user

import (
	"blog-api/rest/req"
	"blog-api/store"
	"blog-api/tools/tokenmanager"
)

type User interface {
	Signup(ctx *req.Ctx) error
	Signin(ctx *req.Ctx) error
	Check(ctx *req.Ctx) error
	Refresh(ctx *req.Ctx) error
	GetByUsername(ctx *req.Ctx) error
}

type Service struct {
	userStore    store.UserStore
	sessionStore store.SessionStore
	tokenManager tokenmanager.Tool
}

func New(userStore store.UserStore, sessionStore store.SessionStore, token tokenmanager.Tool) *Service {
	return &Service{
		userStore,
		sessionStore,
		token,
	}
}
