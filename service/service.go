package service

import (
	"blog-api/service/article"
	"blog-api/service/user"
	"blog-api/store"
	"blog-api/tools/tokenmanager"
)

type Service struct {
	User    user.User
	Article article.Article
}

type ServiceOpts struct {
	UserStore    store.UserStore
	SessionStore store.SessionStore
	ArticleStore store.ArticleStore
	Token        tokenmanager.Tool
	AssetsDir    string
}

func New(s ServiceOpts) *Service {
	return &Service{
		User:    user.New(s.UserStore, s.SessionStore, s.Token),
		Article: article.New(s.UserStore, s.ArticleStore, s.Token, s.AssetsDir),
	}
}
