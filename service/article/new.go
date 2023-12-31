package article

import (
	"blog-api/pkg/errs"
	"blog-api/rest/req"
	"blog-api/store"
	"blog-api/tools/chars"
	"blog-api/tools/tokenmanager"
)

type NewArticleRequest struct {
	Title    string   `json:"title"`
	Markdown string   `json:"markdown"`
	Tags     []string `json:"tags"`
	Preview  string   `json:"preview"`
}

type NewArticleResponse struct {
	Path string `json:"path"`
}

func (s *Service) NewArticle(ctx *req.Ctx) error {
	var r NewArticleRequest

	err := ctx.ParseJSON(&r)
	if err != nil {
		return errs.ReturnError(ctx.Writer, errs.InvalidJSON)
	}

	claims, err := s.tokenManager.ValidateJWTToken(ctx.BearerToken())
	if err != nil {
		return err
	}

	pathSalt, err := tokenmanager.GenerateRandomSalt(2)
	if err != nil {
		return err
	}

	path := chars.ToLatin(r.Title) + "-" + pathSalt

	err = s.articleStore.CreateArticle(store.CreateArticleOpts{
		Title:    r.Title,
		Path:     path,
		Markdown: r.Markdown,
		Tags:     r.Tags,
		Preview:  r.Preview,
		AuthorID: claims.UserID,
	})
	if err != nil {
		return err
	}

	return ctx.JSON(NewArticleResponse{Path: path})
}
