package article

import (
	"blog-api/rest/req"
	"blog-api/store"
	"blog-api/tools/tokenmanager"
)

type Article interface {
	UploadImage(ctx *req.Ctx) error
	SendFile(ctx *req.Ctx) error
	GetArticle(ctx *req.Ctx) error
	NewArticle(ctx *req.Ctx) error
}

type Service struct {
	userStore    store.UserStore
	articleStore store.ArticleStore
	tokenManager tokenmanager.Tool
	assetsDir    string
}

func New(userStore store.UserStore, articleStore store.ArticleStore, token tokenmanager.Tool, assetsDir string) *Service {
	return &Service{
		userStore,
		articleStore,
		token,
		assetsDir,
	}
}
