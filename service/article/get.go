package article

import (
	"blog-api/dto"
	"blog-api/model"
	"blog-api/rest/req"
	"blog-api/store"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (s *Service) GetArticle(ctx *req.Ctx) error {
	path := chi.URLParam(ctx.Request, "path")

	article, err := s.articleStore.GetArticle(store.GetArticleOpts{
		Path: path,
	})
	if err != nil {
		return err
	}

	user, err := s.userStore.GetUserByID(store.GetUserOpts{ID: article.Author_id})
	if err != nil {
		return err
	}

	return ctx.JSON(dto.NewArticle(model.NewArticle(article), dto.NewUserPublic(model.NewUser(user))))
}

func (s *Service) GetArticleForFeed(ctx *req.Ctx) error {
	page := ctx.Request.URL.Query().Get("page")
	if page == "" {
		page = "0"
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return err
	}

	articles, err := s.articleStore.GetArticleForFeed(store.GetArticleFeed{Page: pageInt})
	if err != nil {
		return err
	}

	var users []store.User
	for _, article := range articles {
		user, err := s.userStore.GetUserByID(store.GetUserOpts{ID: article.Author_id})
		if err != nil {
			return err
		}
		users = append(users, user)
	}

	return ctx.JSON(dto.NewArticles(model.NewArticles(articles), dto.NewUsersPublic(model.NewUsers(users))))
}
