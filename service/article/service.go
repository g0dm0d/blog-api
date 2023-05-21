package article

import (
	"blog-api/rest/req"
	"blog-api/tools/tokenmanager"
)

type Article interface {
	UploadImage(ctx *req.Ctx) error
	SendFile(ctx *req.Ctx) error
}

type Service struct {
	tokenManager tokenmanager.Tool
	assetsDir    string
}

func New(token tokenmanager.Tool, assetsDir string) *Service {
	return &Service{
		token,
		assetsDir,
	}
}
