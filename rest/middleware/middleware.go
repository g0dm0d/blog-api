package middleware

import (
	"blog-api/service"
	"blog-api/tools/tokenmanager"
)

type Middleware struct {
	service      *service.Service
	tokenManager *tokenmanager.Tool
}

func New(s *service.Service, t *tokenmanager.Tool) *Middleware {
	return &Middleware{
		service:      s,
		tokenManager: t,
	}
}
