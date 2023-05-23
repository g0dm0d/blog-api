package article

import (
	"blog-api/dto"
	"blog-api/model"
	"blog-api/rest/req"
	"blog-api/store"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type GetArticleResponse struct {
	Name     string         `json:"name"`
	Markdown string         `json:"markdown"`
	Tags     []string       `json:"tags"`
	Preview  string         `json:"preview"`
	Author   dto.UserPublic `json:"author"`
}

func (s *Service) GetArticle(ctx *req.Ctx) error {
	idStr := chi.URLParam(ctx.Request, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	article, err := s.articleStore.GetArticle(store.GetArticleOpts{
		ID: id,
	})
	if err != nil {
		return err
	}

	user, err := s.userStore.GetUserByID(article.Author_id)
	if err != nil {
		return err
	}

	return ctx.JSON(GetArticleResponse{
		Name:     article.Name,
		Markdown: article.Markdown,
		Tags:     article.Tags,
		Preview:  article.Preview,
		Author:   dto.NewUserPublic(model.NewUser(user)),
	})
}
