package article

import (
	"blog-api/dto"
	"blog-api/model"
	"blog-api/rest/req"
	"blog-api/store"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

type GetArticleResponse struct {
	Title     string         `json:"title"`
	Markdown  string         `json:"markdown"`
	Tags      []string       `json:"tags"`
	Preview   string         `json:"preview"`
	Create_at time.Time      `json:"created_at"`
	Author    dto.UserPublic `json:"author"`
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

	user, err := s.userStore.GetUserByID(store.GetUserOpts{ID: article.Author_id})
	if err != nil {
		return err
	}

	return ctx.JSON(GetArticleResponse{
		Title:     article.Title,
		Markdown:  article.Markdown,
		Tags:      article.Tags,
		Preview:   article.Preview,
		Create_at: article.Created_at,
		Author:    dto.NewUserPublic(model.NewUser(user)),
	})
}
