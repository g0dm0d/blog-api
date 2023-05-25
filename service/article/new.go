package article

import (
	"blog-api/pkg/errs"
	"blog-api/rest/req"
	"blog-api/store"
)

type NewArticleRequest struct {
	Title    string   `json:"title"`
	Markdown string   `json:"markdown"`
	Tags     []string `json:"tags"`
	Preview  string   `json:"preview"`
}

func (s *Service) NewArticle(ctx *req.Ctx) error {
	var r NewArticleRequest

	err := ctx.ParseJSON(&r)
	if err != nil {
		errs.ReturnError(ctx.Writer, errs.InvalidJSON)
		return nil
	}

	claims, err := s.tokenManager.ValidateJWTToken(ctx.BearerToken())
	if err != nil {
		return err
	}

	err = s.articleStore.CreateArticle(store.CreateArticleOpts{
		Title:    r.Title,
		Markdown: r.Markdown,
		Tags:     r.Tags,
		Preview:  r.Preview,
		AuthorID: claims.UserID,
	})
	if err != nil {
		return err
	}

	return nil
}
